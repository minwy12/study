# Make


###1. 작성 규칙

```bash
<Target>: <Dependency>
    <Recipe>
```
* Target
  * 빌드 대상의 이름
  * 이 rule에서 만들어낼 타겟의 이름 또는 command
* Dependency
  * 빌드 대상이 의존하는 파일이나 또 다른 target
* Recipe
  * 빌드 대상을 만드는 명령
  * 여러 줄을 입력할 수 있으며 tab으로 시작해야 함
#####예시
    ```bash
    hello:
	    @echo "Hello"

    build:
        @go build -o simple simple.go
    
    run:
        @go run simple.go
    
    clean:
        @rm -f simple
    
    all: hello build run clean
    ```

###2. Command prefix
* @: Don't print command
* -: Ignore error
* +: Run even if Make is in ‘don’t execute’ mode

###3. Phony target
* 이름 뜻
  * 가짜의, 허구의
* 배경
  * 예를 들어 clean target은 clean이라는 파일을 만들지 않음
  * clean이라는 파일이 같은 경로에 없으면 make clean을 실행하는데 문제가 없지만 clean이라는 파일이 있으면 실행되지 않음
    * make: `clean' is up to date.
  * 이때 사용하는 것이 .PHONY