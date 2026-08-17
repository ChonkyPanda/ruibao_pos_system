package ui

// import (
// 	"fmt"
// 	"strconv"
// 	"strings"

// 	"ruibao_pos_system/controller"

// 	"fyne.io/fyne/v2"
// 	"fyne.io/fyne/v2/container"
// 	"fyne.io/fyne/v2/dialog"
// 	"fyne.io/fyne/v2/widget"
// )

// func ShowInventory(w fyne.Window, onBack func()) {
// 	allProducts, err := db.GetAllProducts()
// 	if err != nil {
// 		allProducts = []db.Product{}
// 	}
// 	products := allProducts // this is the filtered/displayed list

// 	headers := []string{"Name", "Barcode", "Price", "Stock", "Category"}

// 	table := widget.NewTable(
// 		func() (int, int) {
// 			return len(products) + 1, len(headers)
// 		},
// 		func() fyne.CanvasObject {
// 			return widget.NewLabel("")
// 		},
// 		func(id widget.TableCellID, obj fyne.CanvasObject) {
// 			label := obj.(*widget.Label)
// 			if id.Row == 0 {
// 				label.TextStyle = fyne.TextStyle{Bold: true}
// 				label.SetText(headers[id.Col])
// 				return
// 			}
// 			label.TextStyle = fyne.TextStyle{}
// 			p := products[id.Row-1]
// 			switch id.Col {
// 			case 0:
// 				label.SetText(p.Name)
// 			case 1:
// 				label.SetText(p.Barcode)
// 			case 2:
// 				label.SetText(fmt.Sprintf("$%.2f", p.Price))
// 			case 3:
// 				label.SetText(strconv.Itoa(p.Stock))
// 			case 4:
// 				label.SetText(p.Category)
// 			}
// 		},
// 	)
// 	table.SetColumnWidth(0, 200)
// 	table.SetColumnWidth(1, 150)
// 	table.SetColumnWidth(2, 90)
// 	table.SetColumnWidth(3, 70)
// 	table.SetColumnWidth(4, 150)

// 	refresh := func() {
// 		updated, err := db.GetAllProducts()
// 		if err != nil {
// 			dialog.ShowError(err, w)
// 			return
// 		}
// 		allProducts = updated
// 		products = allProducts
// 		table.Refresh()
// 	}

// 	// --- Search ---
// 	searchEntry := widget.NewEntry()
// 	searchEntry.SetPlaceHolder("Search by name or barcode...")
// 	searchEntry.OnChanged = func(query string) {
// 		if query == "" {
// 			products = allProducts
// 		} else {
// 			var filtered []db.Product
// 			q := strings.ToLower(query)
// 			for _, p := range allProducts {
// 				if strings.Contains(strings.ToLower(p.Name), q) || strings.Contains(strings.ToLower(p.Barcode), q) {
// 					filtered = append(filtered, p)
// 				}
// 			}
// 			products = filtered
// 		}
// 		table.Refresh()
// 	}

// 	// --- Add Product ---
// 	nameEntry := widget.NewEntry()
// 	nameEntry.SetPlaceHolder("Product name")

// 	barcodeEntry := widget.NewEntry()
// 	barcodeEntry.SetPlaceHolder("Barcode")

// 	categoryEntry := widget.NewEntry()
// 	categoryEntry.SetPlaceHolder("Category")

// 	priceEntry := widget.NewEntry()
// 	priceEntry.SetPlaceHolder("Price (e.g. 4.99)")

// 	stockEntry := widget.NewEntry()
// 	stockEntry.SetPlaceHolder("Starting stock")

// 	addButton := widget.NewButton("Add Product", func() {
// 		if nameEntry.Text == "" || barcodeEntry.Text == "" {
// 			dialog.ShowInformation("Missing info", "Please enter at least a name and barcode.", w)
// 			return
// 		}

// 		price, err := strconv.ParseFloat(priceEntry.Text, 32)
// 		if err != nil {
// 			dialog.ShowInformation("Invalid price", "Please enter a valid number for price.", w)
// 			return
// 		}

// 		stock, err := strconv.Atoi(stockEntry.Text)
// 		if err != nil {
// 			dialog.ShowInformation("Invalid stock", "Please enter a whole number for stock.", w)
// 			return
// 		}

// 		if _, err := db.CreateProduct(nameEntry.Text, barcodeEntry.Text, categoryEntry.Text, float64(price), stock); err != nil {
// 			dialog.ShowError(err, w)
// 			return
// 		}

// 		nameEntry.SetText("")
// 		barcodeEntry.SetText("")
// 		categoryEntry.SetText("")
// 		priceEntry.SetText("")
// 		stockEntry.SetText("")

// 		refresh()
// 		dialog.ShowInformation("Success", "Product added.", w)
// 	})

// 	addForm := container.NewVBox(
// 		widget.NewLabelWithStyle("Add New Product", fyne.TextAlignLeading, fyne.TextStyle{Bold: true}),
// 		nameEntry,
// 		barcodeEntry,
// 		categoryEntry,
// 		priceEntry,
// 		stockEntry,
// 		addButton,
// 	)

// 	// --- Update Stock (placeholder — will become click-row-to-edit later) ---
// 	updateBarcodeEntry := widget.NewEntry()
// 	updateBarcodeEntry.SetPlaceHolder("Barcode")

// 	updateStockEntry := widget.NewEntry()
// 	updateStockEntry.SetPlaceHolder("New stock")

// 	updateButton := widget.NewButton("Update Stock", func() {
// 		if updateBarcodeEntry.Text == "" {
// 			dialog.ShowInformation("Missing barcode", "Please enter a barcode.", w)
// 			return
// 		}

// 		newStock, err := strconv.Atoi(updateStockEntry.Text)
// 		if err != nil {
// 			dialog.ShowInformation("Invalid stock", "Please enter a whole number.", w)
// 			return
// 		}

// 		if err := db.UpdateStock(updateBarcodeEntry.Text, newStock); err != nil {
// 			dialog.ShowError(err, w)
// 			return
// 		}

// 		updateBarcodeEntry.SetText("")
// 		updateStockEntry.SetText("")
// 		refresh()
// 		dialog.ShowInformation("Success", "Stock updated.", w)
// 	})

// 	updateForm := container.NewVBox(
// 		widget.NewLabelWithStyle("Update Stock", fyne.TextAlignLeading, fyne.TextStyle{Bold: true}),
// 		updateBarcodeEntry,
// 		updateStockEntry,
// 		updateButton,
// 	)

// 	deleteBarcodeEntry := widget.NewEntry()
// 	deleteBarcodeEntry.SetPlaceHolder("Barcode")

// 	deleteButton := widget.NewButton("Delete Product", func() {
// 		if deleteBarcodeEntry.Text == "" {
// 			dialog.ShowInformation("Missing barcode", "Please enter a barcode.", w)
// 			return
// 		}

// 		dialog.ShowConfirm("Delete product?", "This cannot be undone from this screen. Continue?", func(confirmed bool) {
// 			if !confirmed {
// 				return
// 			}
// 			if err := db.DeleteProduct(deleteBarcodeEntry.Text); err != nil {
// 				dialog.ShowError(err, w)
// 				return
// 			}
// 			deleteBarcodeEntry.SetText("")
// 			refresh()
// 			dialog.ShowInformation("Deleted", "Product removed.", w)
// 		}, w)
// 	})

// 	deleteForm := container.NewVBox(
// 		widget.NewLabelWithStyle("Delete Product", fyne.TextAlignLeading, fyne.TextStyle{Bold: true}),
// 		deleteBarcodeEntry,
// 		deleteButton,
// 	)

// 	backButton := widget.NewButton("Back", func() {
// 		if onBack != nil {
// 			onBack()
// 		}
// 	})

// 	sidebar := container.NewVBox(
// 		addForm,
// 		widget.NewSeparator(),
// 		updateForm,
// 		widget.NewSeparator(),
// 		deleteForm,
// 		widget.NewSeparator(),
// 		backButton,
// 	)

// 	content := container.NewBorder(
// 		container.NewVBox(
// 			widget.NewLabelWithStyle("Inventory", fyne.TextAlignCenter, fyne.TextStyle{Bold: true}),
// 			searchEntry,
// 		),
// 		nil, nil, sidebar,
// 		table,
// 	)

// 	w.SetContent(content)
// }