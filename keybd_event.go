// Package keybd_event is used for a key press simulated in Windows, Linux and Mac
package keybd_event

type KeyBinding struct {
	hasCTRL   bool
	hasALT    bool
	hasSHIFT  bool
	hasRCTRL  bool
	hasRSHIFT bool
	keys      []int
}

//Use for create struct KeyBinding
func NewKeyBinding() (KeyBinding, error) {
	keyBinding := KeyBinding{}
	keyBinding.hasALT = false
	keyBinding.hasCTRL = false
	keyBinding.hasSHIFT = false
	keyBinding.hasRCTRL = false
	keyBinding.hasRSHIFT = false
	keyBinding.keys = []int{}
	err := initKeyBD()
	if err != nil {
		return keyBinding, err
	}
	return keyBinding, nil
}
func (k *KeyBinding) SetKeys(keys ...int) {
	k.keys = keys
}

//If key ALT pressed
func (k *KeyBinding) HasALT(b bool) {
	k.hasALT = b
}

//If key CTRL pressed
func (k *KeyBinding) HasCTRL(b bool) {
	k.hasCTRL = b
}

//If key SHIFT pressed
func (k *KeyBinding) HasSHIFT(b bool) {
	k.hasSHIFT = b
}
