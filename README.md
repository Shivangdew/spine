# 🦴 Spine
<p align="center">
  <img src="assets/spine-logo.png" alt="Spine" width="420" />
</p>

**Spine is the execution backbone of a backend system.**

Spine defines **how a request is resolved, executed, and completed** — explicitly.

**Spine는 요청 실행(Runtime)을 1급 개념으로 다루는 백엔드 웹 프레임워크입니다.**



## Spine의 문제의식은 단순합니다

대부분의 웹 프레임워크는 다음을 숨깁니다.
요청이 어디서 시작되는지, 누가 인자를 만들고, 언제 비즈니스 코드가 실행되며, 결과가 어떻게 응답으로 변환되는지
Spine은 이 흐름을 숨기지 의도적으로 숨기지 않습니다.


## 한국과 Spine

한국에는 자체 IoC + 실행 파이프라인 구조를 가진 백엔드 프레임워크가 거의 없습니다. 전자정부 프레임워크조차 Spring IoC 위의 조합물입니다. NestJS, Spring, FastAPI, Django는 모두 해외 설계 철학의 수입입니다.

Spine는 한국에서 거의 처음 시도되는, Execution Pipeline 중심의 현대적인 백엔드 웹 프레임워크입니다.

## Spine으로 만든 예제 프로젝트 확인하기
[User-Demo 프로젝트 확인하기](https://github.com/NARUBROWN/spine-user-demo)

## Spine은 무엇이 아닌가

- ❌ HTTP Engine이 아닙니다.
- ❌ Router 중심 프레임워크가 아닙니다.
- ❌ Annotation 기반 자동 프레임워크가 아닙니다.
- ❌ Controller에 책임을 몰아넣지 않습니다.
- ❌ Convention over Configuration(관례 우선)을 채택하지 않습니다.

Spine은 **Execution Pipeline**입니다.


## 전체 아키텍처 개요

```
HTTP Engine (Echo)
        │
        ▼
core.Context
        │
        ▼
Pipeline
  ├─ Router
  ├─ ArgumentResolver 체인
  ├─ Interceptor (preHandle) (구현 예정)
  ├─ Invoker (Method Invocation)
  ├─ ReturnValueHandler
  └─ Interceptor (postHandle) (구현 예정)
        │
        ▼
Response
```
이 흐름은 문서가 아니라 코드로 고정되어 있습니다.

## Execution Pipeline (핵심 모델)

모든 요청은 아래의 실행 순서를 따릅니다.

1. Pipeline 진입
2. Router를 통해 HandlerMethod 선택
3. ArgumentResolver 체인 실행
4. Interceptor.preHandle
5. Controller Method 호출 (Invoker)
6. ReturnValueHandler 실행
7. Interceptor.postHandle
8. Response 생성

이 순서는 숨겨지지 않고, 암묵적으로 바뀌지 않으며, 변경 시 반드시 명시적으로 표현됩니다.

## Controller 철학 (Minimal Responsibility)

Controller는 다음 책임을 가지지 않도록 설계되었습니다.

- HTTP Status 결정
- Header 조작
- Request Parsing
- Argument 생성 규칙
- Response 직렬화

Controller의 책임은 유즈케이스 표현 하나뿐입니다.

```go
func (c *UserController) GetUser(id int) User
```

프레임워크를 모르도록 설계되었으며, 테스트 가능한 순수 구조입니다. 그리고 시그니처 자체가 API 계약입니다.

## Signature-as-Contract

Spine에서 API는 Annotation이 아니라 시그니처입니다.

- 입력 생성 → `ArgumentResolver`
- 출력 표현 → `ReturnValueHandler`

시그니처 변경 = API 변경

Spine는 다음을 의도적으로 금지하도록 설계되었습니다.

- ❌ Annotation 기반 매핑
- ❌ Convention over Configuration (관례 우선)
- ❌ 암묵적 파라미터 주입

## Pipeline과 Invoker의 분리

### Pipeline
- 요청 실행의 전체 흐름을 관리하는 유일한 오케스트레이터.
- 실행 순서를 아는 유일한 컴포넌트입니다.
- 비즈니스 로직을 절대 포함하지 않습니다.

### Invoker
- Controller 인스턴스 생성 (IoC)
- Reflection기반 Method 호출
- Argument / Return 처리의 경계

실행 흐름 제어와 호출 책임을 분리합니다.

## 확장 포인트 (Explicit Extension)

Spine의 모든 확장은 명시적 인터페이스로만 이루어집니다.

### ArgumentResolver
- 메서드 파라미터 하나를 책임집니다.
- Path / Query / Body / DTO 해석 담당
- 모호하면 실패하도록 설계되었습니다.

### ReturnValueHandler
- 반환값 → Response 변환
- JSON / String / Error 등 명확한 책임으로 나눠져있습니다.

### Interceptor (개발 예정)
- 인증, 로깅, 트랜잭션 같은 횡단 관심사 처리
- 실행 흐름에만 관여

> 등록되지 않으면 실행되지 않습니다.

### Container 책임

- Constructor 등록
- Singleton 캐시
- Lazy 생성
- 순환 의존성 감지

> DI는 문법이 아니라 생성 통제 + 그래프 해석입니다.

## Echo와 Spine의 관계

Spine에서 Echo는 HTTP Transport 구현체일 뿐입니다.
Spine 내부 흐름은 다음과 같습니다.
```
Echo → core.Context → Spine Runtime
```
Echo 타입은 Spine 내부에 노출되지 않습니다.
또한, 교체 가능합니다.

## License

MIT

## Status

✅ 이미 개발 완료
- Execution Pipeline 구조 확정 
- Router + HandlerMeta 구현
- Invoker (Reflection 기반 메서드 실행)
- ArgumentResolver 체계 구축
- ContextResolver (core.Context 주입)
- PrimitiveResolver (Path 1개 / Query 1개 자동 매핑)
- QueryDTOResolver (query 태그 기반)
- Body DTOResolver (JSON Body 바인딩)
- Resolver Registry + 우선순위 체계
- ReturnValueHandler (JSON / String / Error)
- ReturnHandler Registry
- IoC Container (Constructor 등록, Lazy 생성)
- 순환 의존성 감지
- Echo Adapter (단일 /* 엔트리 포인트)
- core.Context 분리 및 Request/Response 계약
- Controller / Service / Repository / Route 분리 예제 제작
- Path + QueryDTO 혼합 사용 가능

🟡 개발 중
- PathDTOResolver 구현 (path:"id" 태그 기반)
- Error 반환 → HTTP Status 매핑 규칙 정리
- Interceptor 구현
- Resolver / Handler 에러 메시지 통일

🟠 개발 예정
- Validation 태그 지원
- Default 값 처리
- Pagination QueryDTO 패턴
- 테스트 유틸리티 제공 (Invoker / Resolver 단위 테스트)

❌ 개발 예정 없음
- Annotation / Decorator 기반 설계
- Component Scan
- Convention over Configuration
- Controller Interface 강제
- 암묵적 / 순서 기반 파라미터 매핑
