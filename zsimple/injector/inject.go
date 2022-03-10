//go:build wireinject
// +build wireinject

package injector

import (
	"io"
	"os"
	"pzn-restful-api/zsimple"

	"github.com/google/wire"
)

// buat function injectornya, tidak ada ketentuan penamaan. Return valuenya, adalah class terakhir yang membutuhkan dependency (service butuh repository)
func InitializedService(isError bool) (*zsimple.SimpleService, error){
	// sebutkan provider mana saja yang akan kita gunakan. Nantinya, siapa yang butuh siapa akan ditangani oleh google wirenya.
	wire.Build(zsimple.NewSimpleRepository, zsimple.NewSimpleService)
	// Return nil karena isi dari function ini nanti hasil compilenya akan diganti oleh google wire setelah kita menyebutkan dependency / provicer apa saja yang ingin kita tambahkan
	return nil, nil
}

func InitializedDBRepo() *zsimple.DatabaseRepository{
	// daftarkan provider yang akan digunakan.
	wire.Build(
		zsimple.NewDatabasePostgre, 
		zsimple.NewDatabaseMongoDB,
		zsimple.NewDatabaseRepository,
	)
	
	return nil
}

var fooSet = wire.NewSet(zsimple.NewFooRepository, zsimple.NewFooService)
var barSet = wire.NewSet(zsimple.NewBarRepository, zsimple.NewBarService)

func InitializedFooBarService() *zsimple.FooBarService{
	// daftarkan provider yang akan digunakan.
	wire.Build(
		fooSet,
		barSet,
		zsimple.NewFooBarService,
	)
	return nil
}

//binding interface, memberitahu bahwa jika butuh parameter interface SayHello, maka kirimkan SayHelloImpl
var helloSet = wire.NewSet(
	zsimple.NewSayHelloImpl,
	wire.Bind(new(zsimple.SayHello), new(*zsimple.SayHelloImpl)),
)

func InitializedHelloService() *zsimple.HelloService{
	wire.Build(
		helloSet, zsimple.NewHelloService,
	)
	return nil
}

// build dengan inject data Foo dan Bar ke sebuah pointer struct baru
func InitializedFooBar() *zsimple.FooBar{
	wire.Build(
		// inject ke semua field (Foo dan Bar) milik struct FooBar
		zsimple.NewFoo, zsimple.NewBar, wire.Struct(new(zsimple.FooBar), "*"),
	)
	return nil
}

var fooValue = &zsimple.Foo{}
var barValue = &zsimple.Bar{}

func InitializedFooBarUsingValue() *zsimple.FooBar{
	wire.Build(
		// sebelumnya kita gunakan function provider NewFoo dan NewBar sebagai value yang akan diinject ke FooBar
		// zsimple.NewFoo, zsimple.NewBar, wire.Struct(new(zsimple.FooBar), "*"),
		// kita bisa langsung sebutkan value apa mau diinject ke struct apa
		wire.Value(fooValue), 
		wire.Value(barValue), 
		wire.Struct(new(zsimple.FooBar),"*"),
	)
	return nil
}

func InitializedReader() io.Reader{
	// jika ada yang butuh data dengan tipe interface io.Reader, inject os.Stdin sebagai valuenya
	wire.Build(wire.InterfaceValue(new(io.Reader), os.Stdin))
	return nil
}

func InitializedConfiguration() *zsimple.Configuration{
	wire.Build(
		zsimple.NewApplication,
		// buat sebuah provider dari field Application, yang nama fieldnya adalah configuration
		// Setiap kali butuh data Configuration, akan diambil dari field configuration dari objek Application
		wire.FieldsOf(new(*zsimple.Application), "Configuration"),
	)
	return nil
}

func InitializedConnection(name string) (*zsimple.Connection, func()){
	wire.Build(
		zsimple.NewConnection,
		zsimple.NewFile,
	)
	return nil, nil
}