package types

type ClientParams struct {
	File   string `short:"f" long:"file" description:"File to transfer to server" value-name:"FILE" required:"true"`
	Server string `short:"s" long:"server" description:"Server address" value-name:"SERVER" required:"true"`
}

type ServerParams struct {
	File string `short:"f" long:"file" description:"File to transfer to server" value-name:"FILE" required:"true"`
	Port int    `short:"p" long:"port" description:"Port to listen" value-name:"PORT" required:"true"`
}

type TestParams struct {
}

