package model

import (
	"github.com/google/uuid"
	_ "gorm.io/gorm"
)

type Task struct {
	ID       uuid.UUID
	Name     string
	Finished bool
}

func GetTasks() ([]Task, error) {
	var tasks []Task
	err := db.Find(&tasks).Error

	return tasks, err
}

func AddTask(name string) (*Task, error) {
	id, err := uuid.NewUUID()
	if err != nil {
		return nil, err
	}

	task := Task{
		ID:       id,
		Name:     name,
		Finished: false,
	}

	// taskをDBのTaskテーブルに追加。その成否を(ry
	err = db.Create(&task).Error

	// taskのポインタ と errを返す
	return &task, err
}

// 関数 ChangeFinishedTaskの引数はuuid.UUID型のtaskIDで、戻り値はerror型である
func ChangeFinishedTask(taskID uuid.UUID) error {

	// DBのTaskテーブルからtaskIDと一致するidを探し、そのFinishedをtureにする(*3)
	err := db.Model(&Task{}).Where("id = ?", taskID).Update("finished", true).Error
	return err
}

func DeleteTask(taskID uuid.UUID) error {
	// DBのTaskテーブルからtaskIDと一致するidを探し、そのタスクを削除する
	err := db.Where("id = ?", taskID).Delete(&Task{}).Error
	return err
}
