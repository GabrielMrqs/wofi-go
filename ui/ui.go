package ui

import (
	"fmt"
	"os/exec"

	"github.com/gotk3/gotk3/glib"
	"github.com/gotk3/gotk3/gtk"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func InitUI(apps map[string]string) {
	// Inicializa GTK
	gtk.Init(nil)

	// Cria janela principal
	win, err := gtk.WindowNew(gtk.WINDOW_TOPLEVEL)
	check(err)
	win.SetTitle("Launcher")
	win.SetDefaultSize(400, 300)
	win.Connect("destroy", func() {
		gtk.MainQuit()
	})

	// Modelo de lista
	listStore, err := gtk.ListStoreNew(glib.TYPE_STRING)
	check(err)

	for name := range apps {
		iter := listStore.Append()
		listStore.SetValue(iter, 0, name)
	}

	treeView, err := gtk.TreeViewNewWithModel(listStore)
	check(err)

	// Coluna de texto
	renderer, err := gtk.CellRendererTextNew()
	check(err)
	column, err := gtk.TreeViewColumnNewWithAttribute("Aplicativo", renderer, "text", 0)
	check(err)
	treeView.AppendColumn(column)

	// Clique na linha
	treeView.Connect("row-activated", func(tv *gtk.TreeView, path *gtk.TreePath, col *gtk.TreeViewColumn) {
		iter, _ := listStore.GetIter(path)
		val, _ := listStore.GetValue(iter, 0)
		appName, _ := val.GetString()
		cmd := apps[appName]
		fmt.Println("Executando:", cmd)

		// Executa o comando via shell
		go func() {
			_ = glib.IdleAdd(func() bool {
				err := exec.Command("sh", "-c", cmd).Start()
				if err != nil {
					fmt.Println("Erro ao executar:", err)
				}
				return false
			})
		}()
	})

	// Scrolled window
	scroll, _ := gtk.ScrolledWindowNew(nil, nil)
	scroll.SetPolicy(gtk.POLICY_AUTOMATIC, gtk.POLICY_AUTOMATIC)
	scroll.Add(treeView)

	// Adiciona ao container
	win.Add(scroll)
	win.ShowAll()

	gtk.Main()
}
