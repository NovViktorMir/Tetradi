package main



func main() {
	println ("hello world")
	var ii int = 25                                                       // var название до int8 16 32 64 типы переменых по битности разный диопозон значсений
	var autoInt = -22                                                     //  autoInt автомотическое распознание int переменной
	var bigInt int64 = 100500                                             //устанавливаем тип перемнной (big)intZZ
	var unsignedInt uint = 9280                                           // устанавливает переменную uint только с положительным значением
	var unsignedbigInt uint64 = 1111111                                   // сттанаваливает переменную uint определёной битности с положительным значением
	println("integers", ii, autoInt, bigInt, unsignedInt, unsignedbigInt) // выводим в командую строку переменные в ()
	var f32 float32 = 2.222                                               //типы с плавующей точкой float32 64
	var f64 float64 = 3.3333
	println("faloat:", f32, f64)
	var boolPav bool = true // имя xxxx bool= булевое значение правда лож да нет пропустить или не пропустить потвердить или не потвердить
	println("flag is", boolPav)
	var hello string = "go\n\t" // спец символ \n новая строка \t подобие отступа // строки
	var world = "world"
	println(hello, world)
    var defaultInt int //переменные по умолчанию
    var defaultFlot float32
    var defaultBool bool
    println("default valueas", defaultBool, defaultFlot, defaultInt)
    var rawSymbol byte // не разобрался с применением ASCIL
    println("symbol", string(rawSymbol))
    shortDeclare :=66
    println("shot declare", shortDeclare)
	//приведение типов
	println("float to int", int(f32))
	println("int to string", string(`1`))// приводим стринг из ASKIL по порядковому номеру 52=4, нужно разобраться в strcony? Типизацыя цыфр и строк
	//комплексные числа
	z := 2+3i
	println("complex number", z)
	v1 := "Viktor "
	v2 := "Gorobsov"
	fullName := v1+v2
	println("My name is", fullName, len(fullName))
	//многострочное обьявление
	escaping := `hello\r\n
	world`
	println("<pre>", escaping, "</pre>")
   // несколько переменных
   var b1, b2 string = "v1", "v2"
   println(b1,b2)
   var (
	m0 int = 12
	m2     = "string"
	m3     = 23
   )

   println(m0, m2, m3)
   // выполнить код до этого места
}