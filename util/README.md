Golang에서 MVC 패턴을 구현할 때, 공통으로 사용되는 함수는 보통 "util" 디렉토리나 "helper" 디렉토리와 같은 곳에 위치시킵니다.

이 디렉토리에는 컨트롤러, 모델 등의 코드와는 별도로 일반적인 기능을 제공하는 패키지를 두는 것이 일반적입니다. 이러한 유틸리티 함수는 모든 컨트롤러와 모델에서 사용할 수 있으며, 코드의 중복을 줄이는 데 도움이 됩니다.

일반적으로 이러한 유틸리티 함수는 구조체를 사용하지 않는 함수로 작성됩니다. 이는 유틸리티 함수가 다양한 구조체에서 사용될 수 있으며, 따라서 구조체와는 관련이 없는 범용 함수로 작성하는 것이 가장 좋기 때문입니다
