package frames
# Remove this line and replace NAME with the name as in: localhost:8080/NAME
func init() {
	FrameMap["NAME"] = DefaultFrameType(_NAME)
}

var _NAME = []string{
	`Frame1 Line1
Frame1 Line2`,

	`Frame2 Line1
Frame2 Line2`,

	`Frame3 Line1
Frame3 Line2`,
}
