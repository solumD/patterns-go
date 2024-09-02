package adapter

func AdapterExample() {
    
    client := &Client{}
    mac := &Mac{}

    client.InsertLightningConnectorIntoComputer(mac)

    windowsPC := &Windows{}
    windowsMachineAdapter := &WindowsAdapter{
        windowsComputer: windowsPC,
    }

    client.InsertLightningConnectorIntoComputer(windowsMachineAdapter)
}