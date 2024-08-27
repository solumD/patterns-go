package abstractfactory

func AbstractFactoryExample() {
	factories := []Factory{}
	companies := []string{"Ozon", "Yandex"}

	for _, c := range companies {
		switch c {
		case "Ozon":
			f := NewOzonFactory()
			factories = append(factories, f)
		case "Yandex":
			f := NewYandexFactory()
			factories = append(factories, f)
		}
	}

	for _, f := range factories {
		a := f.GetApp(100, 100)
		a.PrintGain()
		a.PrintUsers()

		s := f.GetStaff(100)
		s.PrintQuantity()
	}
}
