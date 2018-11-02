package entity

//stroage
import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"
)

// UserFilter : UserFilter types take an *User and return a bool value.
type UserFilter func(*User) bool

// MeetingFilter : MeetingFilter types take an *User and return a bool value.
type MeetingFilter func(*Meeting) bool

var userinfoPath = "./userInfo.txt"
var metinfoPath = "./meetingInfo.txt"
var curUserPath = "./curUser.txt"

var dirty bool

var uData []User
var mData []Meeting

var curUser User

func init() {
	dirty = false
	if err := readFromFile(); err != nil {
		fmt.Println("readFromFile fail:", err)
	}
}

// Logout : logout
func Logout() error {
	curUser.Username = ""
	return Sync()
}

// Sync : sync file
func Sync() error {
	if err := writeToFile(); err != nil {
		fmt.Println("writeToFile fail:", err)
		return err
	}
	return nil
}

// CreateUser : create a user
// @param a user object
func CreateUser(v *User) {
	uData = append(uData, (*v))
	dirty = true
}

// QueryUser : query users
// @param a lambda function as the filter
// @return a list of fitted users
func QueryUser(filter UserFilter) []User {
	var user []User
	for _, v := range uData {
		if filter(&v) {
			user = append(user, v)
		}
	}
	return user
}

// UpdateUser : update users
// @param a lambda function as the filter
// @param a lambda function as the method to update the user
// @return the number of updated users
func UpdateUser(filter UserFilter, switcher func(*User)) int {
	count := 0
	for i := 0; i < len(uData); i++ {
		if v := &uData[i]; filter(v) {
			switcher(v)
			count++
		}
	}
	if count > 0 {
		dirty = true
	}
	return count
}

// DeleteUser : delete users
// @param a lambda function as the filter
// @return the number of deleted users
func DeleteUserD(filter UserFilter) int {
	count := 0
	length := len(uData)
	for i := 0; i < length; {
		if filter(&uData[i]) {
			length--
			uData[i] = uData[length]
			uData = uData[:length]
			count++
		} else {
			i++
		}
	}
	if count > 0 {
		dirty = true
	}
	return count
}

// CreateMeeting : create a meeting
// @param a meeting object
func CreateMeetingD(v *Meeting) {
	mData = append(mData, *v)
	dirty = true
}

// QueryMeeting : query meetings
// @param a lambda function as the filter
// @return a list of fitted meetings
func QueryMeetingD(filter MeetingFilter) []Meeting {
	var met []Meeting
	for _, v := range mData {
		if filter(&v) {
			met = append(met, v)
		}
	}
	return met
}

// UpdateMeeting : update meetings
// @param a lambda function as the filter
// @param a lambda function as the method to update the meeting
// @return the number of updated meetings
func UpdateMeeting(filter MeetingFilter, switcher func(*Meeting)) int {
	count := 0
	for i := 0; i < len(mData); i++ {
		if v := &mData[i]; filter(v) {
			switcher(v)
			count++
		}
	}
	if count > 0 {
		dirty = true
	}
	return count
}

// DeleteMeeting : delete meetings
// @param a lambda function as the filter
// @return the number of deleted meetings
func DeleteMeetingD(filter MeetingFilter) int {
	count := 0
	length := len(mData)
	for i := 0; i < length; {
		if filter(&mData[i]) {
			length--
			mData[i] = mData[length]
			mData = mData[:length]
			count++
		} else {
			i++
		}
	}
	if count > 0 {
		dirty = true
	}
	return count
}

// GetCurUser : get current user
// @return the current user
// @return error if current user does not exist
func GetCurUserD() (User, error) {
	if curUser.Username == "" {
		return User{}, errors.New("Current user does not exist")
	}
	return curUser, nil

}

// SetCurUser : get current user
// @param current user
func SetCurUser(u *User) {
	curUser.Username = u.Username
}

// readFromFile : read file content into memory
// @return if fail, error will be returned
func readFromFile() error {
	var e []error
	err1 := readJSON(curUserPath, curUser)
	if err1 != nil {
		e = append(e, err1)
	}

	if err := readUser(); err != nil {
		e = append(e, err)
	}
	if err := readMet(); err != nil {
		e = append(e, err)
	}
	if len(e) == 0 {
		return nil
	}
	er := e[0]
	for i := 1; i < len(e); i++ {
		er = errors.New(er.Error() + e[i].Error())
	}
	return er
}

// writeToFile : write file content from memory
// @return if fail, error will be returned
func writeToFile() error {
	var e []error
	if err := writeJSON(curUserPath, curUser); err != nil {
		e = append(e, err)
	}
	if dirty {
		if err := writeJSON(userinfoPath, uData); err != nil {
			e = append(e, err)
		}
		if err := writeJSON(metinfoPath, mData); err != nil {
			e = append(e, err)
		}
	}
	if len(e) == 0 {
		return nil
	}
	er := e[0]
	for i := 1; i < len(e); i++ {
		er = errors.New(er.Error() + e[i].Error())
	}
	return er
}

func readUser() error {
	return readJSON(userinfoPath, uData)
}

func readMet() error {
	return readJSON(metinfoPath, mData)
}

func readJSON(fpath string, data interface{}) error {
	file, err := os.Open(fpath)
	if err != nil {
		fmt.Println("Open File Fail:", fpath, err)
		return err
	}
	defer file.Close()

	dec := json.NewDecoder(file)
	switch err := dec.Decode(&data); err {
	case nil, io.EOF:
		return nil
	default:
		fmt.Println("Decode User Fail:", err)
		return err
	}
}

func writeJSON(fpath string, data interface{}) error {
	file, err := os.Create(fpath)
	if err != nil {
		return err
	}
	defer file.Close()

	enc := json.NewEncoder(file)

	if err := enc.Encode(&data); err != nil {
		fmt.Println("writeJSON:", err)
		return err
	}
	return nil
}
