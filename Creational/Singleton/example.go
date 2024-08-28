package singleton

func SingletonExample() {
	singleton1 := InitSingleton("First")
	singleton1.PrintName() // First
	singleton2 := InitSingleton("Second")
	singleton2.PrintName() // First
}
