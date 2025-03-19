package exports

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/OdairPianta/julia/models"
	"github.com/xuri/excelize/v2"
)

func ExportSampleModels(models []models.SampleModel) (string, error) {
	file := excelize.NewFile()
	defer func() {
		if err := file.Close(); err != nil {
			return
		}
	}()
	index, err := file.NewSheet("Sheet1")
	if err != nil {
		return "", err
	}
	file.SetActiveSheet(index)
	file.SetCellValue("Sheet1", "A1", "Nome")
	file.SetCellValue("Sheet1", "B1", "Exemplo String")
	file.SetCellValue("Sheet1", "C1", "Exemplo Único")
	file.SetCellValue("Sheet1", "D1", "Exemplo Date")
	file.SetCellValue("Sheet1", "E1", "Exemplo Anulável")
	file.SetCellValue("Sheet1", "F1", "Exemplo Decimal")
	file.SetCellValue("Sheet1", "G1", "Exemplo Detalhe")

	for i, model := range models {
		file.SetCellValue("Sheet1", fmt.Sprintf("A%d", i+2), model.Name)
		file.SetCellValue("Sheet1", fmt.Sprintf("B%d", i+2), model.SampleString)
		file.SetCellValue("Sheet1", fmt.Sprintf("C%d", i+2), model.SampleUnique)
		file.SetCellValue("Sheet1", fmt.Sprintf("D%d", i+2), model.SampleDate.Format("02/01/2006"))
		file.SetCellValue("Sheet1", fmt.Sprintf("E%d", i+2), model.SampleNullable)
		file.SetCellValue("Sheet1", fmt.Sprintf("F%d", i+2), model.SampleDouble)
		if model.SampleDetail != nil {
			file.SetCellValue("Sheet1", fmt.Sprintf("G%d", i+2), model.SampleDetail.SampleString)
		}
	}
	wd, err := os.Getwd()
	if err != nil {
		return "", err
	}
	wd = strings.Replace(wd, "/tests", "", 1)

	tempDir := filepath.Join(wd, "storage", "temp")

	fileName := filepath.Join(tempDir, time.Now().Format("2006_01_02_150405")+"_samples_export.xlsx")
	if err := file.SaveAs(fileName); err != nil {
		return "", err
	}

	return fileName, nil
}
