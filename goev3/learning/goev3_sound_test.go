package learning

import (
	"fmt"
	"os"
	"testing"

	"path/filepath"

	"github.com/mattrajca/GoEV3/Sound"
	"github.com/mattrajca/GoEV3/utilities"
)

type MockFileReader struct {
	defaultFileReader utilities.DefaultFileReader
}

func (fr MockFileReader) ReadFile(filename string) ([]byte, error) {
	filename = filepath.Join("/tmp", filename)
	fmt.Printf("Reading File %v\n", filename)
	return fr.defaultFileReader.ReadFile(filename)
}

type MockFileWriter struct {
	defaultFileWriter utilities.DefaultFileWriter
}

func (fw MockFileWriter) WriteFile(filename string, data []byte, perm os.FileMode) error {
	filename = filepath.Join("/tmp", filename)
	fmt.Printf("Writing File %v with content %v\n", filename, data)
	return fw.defaultFileWriter.WriteFile(filename, data, perm)
}

func Test_Audio_SetVolumeAndPlayToneShouldKeepTheVolumeAndResetToneAfterPlayed(t *testing.T) {
	// Arange
	utilities.FReader = MockFileReader{}
	utilities.FWriter = MockFileWriter{}

	// Run
	Sound.SetVolume(50)
	Sound.PlayTone(100, 1000)

	// Verify
	actualVolume := Sound.CurrentVolume()
	if actualVolume != 50 {
		t.Errorf("%v: expected %v but was %v", "volume", 50, actualVolume)
	}

	actualTone := Sound.CurrentTone()
	if actualTone != 0 {
		t.Errorf("%v: expected %v but was %v", "volume", 0, actualTone)
	}

}
