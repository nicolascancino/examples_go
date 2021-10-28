package main

import (
	"fmt"
)

type taskList struct {
	tasks []*task
}

func (taskList *taskList) add(task *task) {
	taskList.tasks = append(taskList.tasks, task)
}
func (taskList *taskList) remove(index int) {
	taskList.tasks = append(taskList.tasks[:index], taskList.tasks[index+1:]...)
}
func (taskList *taskList) printList() {
	for _, task := range taskList.tasks {
		if task.completed {
			fmt.Println("Nombre: ", task.name)
			fmt.Println("Descripcion: ", task.description)
			fmt.Println("Completada: ", task.completed)
		}
	}
}

type task struct {
	name        string
	description string
	completed   bool
}

func (task *task) setName(name string) {
	task.name = name
}

func (task *task) setDescription(description string) {
	task.description = description
}
func (task *task) done() {
	task.completed = true
}

func main() {

	task1 := &task{
		name:        "Completar curso go",
		description: "Terminar el curso de go",
	}

	task2 := &task{
		name:        "Completar curso python",
		description: "Terminar el curso de python",
	}

	task3 := &task{
		name:        "Completar curso nodeJs",
		description: "Terminar el curso de nodeJs",
	}

	taskList1 := &taskList{
		tasks: []*task{
			task1,
			task2,
		},
	}

	taskList1.add(task3)
	taskList1.tasks[1].done()
	taskList1.printList()
	taskList1.remove(2)

	for i := 0; i < len(taskList1.tasks); i++ {
		fmt.Println(taskList1.tasks[i])
	}

	for index, task := range taskList1.tasks {
		fmt.Println(index, task)
	}

	mapTasks := make(map[string]*taskList)

	mapTasks["Nicolas"] = taskList1

	fmt.Println("Tareas Nicolas")
	mapTasks["Nicolas"].printList()

}
