Hello {{ .Name }}
Number is {{ .Number | twice }}
Number is {{ .Number | addSomeNumber 8 }}
String is {{ .String | joinStrings "world" }}
String is {{ joinStrings "world" .String }}