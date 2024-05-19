package tui

import "github.com/charmbracelet/bubbles/key"

type navigationKeymap struct {
	Enter key.Binding
	Back  key.Binding
	Quit  key.Binding
	Next  key.Binding
	Prev  key.Binding
}

type imgKeymap struct {
	Create      key.Binding
	Rename      key.Binding
	Pull        key.Binding
	Prune       key.Binding
	Delete      key.Binding
	DeleteForce key.Binding
}

type contKeymap struct {
	ToggleListAll   key.Binding
	ToggleStartStop key.Binding
	Delete          key.Binding
	DeleteForce     key.Binding
	Exec            key.Binding
	Prune           key.Binding
}

type volKeymap struct {
	Delete key.Binding
}

var ImageKeymap = imgKeymap{
	Create: key.NewBinding(
		key.WithKeys("c"),
		key.WithHelp("c", "create"),
	),
	Rename: key.NewBinding(
		key.WithKeys("r"),
		key.WithHelp("r", "rename"),
	),
	Delete: key.NewBinding(
		key.WithKeys("d"),
		key.WithHelp("d", "delete"),
	),
	DeleteForce: key.NewBinding(
		key.WithKeys("D"),
		key.WithHelp("D", "delete (force)"),
	),
	Pull: key.NewBinding(
		key.WithKeys("o"),
		key.WithHelp("o", "Pull new Image"),
	),
	Prune: key.NewBinding(
		key.WithKeys("p"),
		key.WithHelp("p", "Prune images"),
	),
}

var ContainerKeymap = contKeymap{
	ToggleListAll: key.NewBinding(
		key.WithKeys("a"),
		key.WithHelp("a", "Toggle list all"),
	),
	ToggleStartStop: key.NewBinding(
		key.WithKeys("s"),
		key.WithHelp("s", "Toggle Start/Stop"),
	),
	Delete: key.NewBinding(
		key.WithKeys("d"),
		key.WithHelp("d", "delete"),
	),
	DeleteForce: key.NewBinding(
		key.WithKeys("D"),
		key.WithHelp("D", "delete (force)"),
	),
	Prune: key.NewBinding(
		key.WithKeys("p"),
		key.WithHelp("p", "prune"),
	),
	Exec: key.NewBinding(
		key.WithKeys("x"),
		key.WithHelp("x", "exec"),
	),
}

var VolumeKeymap = volKeymap{
	Delete: key.NewBinding(
		key.WithKeys("d"),
		key.WithHelp("d", "delete"),
	),
}

var NavKeymap = navigationKeymap{
	Enter: key.NewBinding(
		key.WithKeys("enter"),
		key.WithHelp("enter", "select"),
	),
	Back: key.NewBinding(
		key.WithKeys("esc"),
		key.WithHelp("esc", "back"),
	),
	Quit: key.NewBinding(
		key.WithKeys("ctrl+c", "q"),
		key.WithHelp("ctrl+c/q", "quit"),
	),
	Next: key.NewBinding(
		key.WithKeys("right", "l", "tab"),
		key.WithHelp("->/l/tab", "next"),
	),
	Prev: key.NewBinding(
		key.WithKeys("left", "h", "shift+tab"),
		key.WithHelp("<-/h/shift+tab", "prev"),
	),
}

func getVolumeKeymap() []key.Binding {
	return []key.Binding{
		VolumeKeymap.Delete,
	}
}

func getImageKeymap() []key.Binding {
	return []key.Binding{
		ImageKeymap.Delete,
		ImageKeymap.DeleteForce,
		ImageKeymap.Prune,
		ImageKeymap.Pull,
	}
}

func getContainerKeymap() []key.Binding {
	return []key.Binding{
		ContainerKeymap.ToggleListAll,
		ContainerKeymap.ToggleStartStop,
		ContainerKeymap.Delete,
		ContainerKeymap.DeleteForce,
		ContainerKeymap.Prune,
		ContainerKeymap.Exec,
	}
}
