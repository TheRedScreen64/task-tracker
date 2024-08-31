package task

import (
	"fmt"
	"strconv"
	"time"

	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/lipgloss/table"
	"github.com/dustin/go-humanize"
)

type TaskStatus string

var (
	TASK_STATUS_TODO        TaskStatus = "todo"
	TASK_STATUS_IN_PROGRESS TaskStatus = "in-progress"
	TASK_STATUS_DONE        TaskStatus = "done"
)

type Task struct {
	ID          uint64     `json:"id"`
	Description string     `json:"description"`
	Status      TaskStatus `json:"status"`
	CreatedAt   time.Time  `json:"createdAt"`
	UpdatedAt   time.Time  `json:"updatedAt"`
}

func ListTasks(filter string) error {
	tasks, err := LoadTasks()
	if err != nil {
		return err
	}

	var statusFilter string
	switch filter {
	case "":
	case string(TASK_STATUS_TODO):
		statusFilter = string(TASK_STATUS_TODO)
	case string(TASK_STATUS_IN_PROGRESS):
		statusFilter = string(TASK_STATUS_IN_PROGRESS)
	case string(TASK_STATUS_DONE):
		statusFilter = string(TASK_STATUS_DONE)
	default:
		return fmt.Errorf(`invalid filter: "%s"`, filter)
	}

	fmt.Println(lipgloss.NewStyle().
		Bold(true).
		Padding(1, 0).
		Foreground(lipgloss.Color("3")).
		Render("Your tasks"))

	headerStyle := lipgloss.NewStyle().Bold(true).Padding(0, 1).Foreground(lipgloss.Color("15"))
	rowStyle := lipgloss.NewStyle().Padding(0, 1)

	t := table.New().
		Border(lipgloss.NormalBorder()).
		BorderStyle(lipgloss.NewStyle().Foreground(lipgloss.Color("7"))).
		Width(100).
		StyleFunc(func(row, col int) lipgloss.Style {
			switch {
			case row == 0:
				if col == 0 {
					return headerStyle.Padding(0)
				}
				return headerStyle
			case col == 0:
				return rowStyle.Padding(0).Align(lipgloss.Right)
			default:
				return rowStyle
			}
		}).
		BorderLeft(false).
		BorderRight(false).
		BorderColumn(false).
		BorderBottom(false).
		Headers("id", "description", "status", "created", "updated")

	for _, task := range *tasks {
		if statusFilter != "" && string(task.Status) != statusFilter {
			continue
		}

		t.Row(
			lipgloss.NewStyle().
				Foreground(lipgloss.Color("8")).
				Render(strconv.FormatUint(task.ID, 10)),
			task.Description,
			lipgloss.NewStyle().
				Foreground(lipgloss.Color(statusColor(task.Status))).
				Render(string(task.Status)),
			humanize.Time(task.CreatedAt),
			humanize.Time(task.UpdatedAt))
	}
	fmt.Print(t, "\n", "\n")

	return nil
}

func AddTask(description string) error {
	tasks, err := LoadTasks()
	if err != nil {
		return err
	}

	taskId := uint64(1)
	if len(*tasks) > 0 {
		lastTaskId := (*tasks)[len(*tasks)-1].ID
		taskId = lastTaskId + 1
	}

	*tasks = append(*tasks, Task{
		ID:          taskId,
		Description: description,
		Status:      TASK_STATUS_TODO,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	})

	if err := SaveTasks(tasks); err != nil {
		return err
	}

	fmt.Println(lipgloss.NewStyle().
		Padding(1, 0).
		Foreground(lipgloss.Color("2")).
		Render("✅ Task was added successfully."))

	return nil
}

func DeleteTask(id uint64) error {
	tasks, err := LoadTasks()
	if err != nil {
		return err
	}

	i, err := getIndexOfTask(tasks, id)
	if err != nil {
		return err
	}

	*tasks = append((*tasks)[:i], (*tasks)[i+1:]...)

	if err := SaveTasks(tasks); err != nil {
		return err
	}

	fmt.Println(lipgloss.NewStyle().Padding(1, 0).Foreground(lipgloss.Color("2")).Render(fmt.Sprintf("✅ Task %v was deleted successfully.", id)))

	return nil
}

func UpdateTaskDesc(id uint64, newDesc string) error {
	tasks, err := LoadTasks()
	if err != nil {
		return err
	}

	i, err := getIndexOfTask(tasks, id)
	if err != nil {
		return err
	}

	(*tasks)[i].Description = newDesc
	(*tasks)[i].UpdatedAt = time.Now()

	if err := SaveTasks(tasks); err != nil {
		return err
	}

	fmt.Println(lipgloss.NewStyle().
		Padding(1, 0).
		Foreground(lipgloss.Color("2")).
		Render("✅ Task description was updated successfully."))

	return nil
}

func UpdateTaskStatus(id uint64, newStatus TaskStatus) error {
	tasks, err := LoadTasks()
	if err != nil {
		return err
	}

	i, err := getIndexOfTask(tasks, id)
	if err != nil {
		return err
	}

	(*tasks)[i].Status = newStatus
	(*tasks)[i].UpdatedAt = time.Now()

	if err := SaveTasks(tasks); err != nil {
		return err
	}

	fmt.Println(lipgloss.NewStyle().
		Padding(1, 0).
		Foreground(lipgloss.Color("2")).
		Render("✅ Task status was updated successfully."))

	return nil
}
