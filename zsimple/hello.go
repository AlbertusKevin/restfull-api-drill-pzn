package zsimple

type SayHello interface {
	Hello(name string) string
}

type SayHelloImpl struct {
}

func (s *SayHelloImpl) Hello(name string) string {
	return "Hello " + name
}

type HelloService struct {
	SayHello SayHello
}

// provider
func NewSayHelloImpl() *SayHelloImpl {
	return &SayHelloImpl{}
}

// provider ini menerima parameter interface SayHello. Namun tidak ada provider yang kita buat yang return interface SayHello ini. Adanya langsung implementasinya. Google wire tetap menganggap bahwa tidak ada provider yang return interface SayHello, meski sebenarnya ada implementasi dari interface tersebut. Keduanya akan dianggap berbeda
func NewHelloService(sayHello SayHello) *HelloService {
	return &HelloService{SayHello: sayHello}
}