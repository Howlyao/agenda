package entity

import (
	"fmt"
)

var curuserinfoPath = "./src/github.com/Howlyao/agenda/entity/curUser.txt"

func UserLogout() bool {
	if err := Logout(); err != nil {
		return false
	} else {
		return true
	}
}
func GetCurUser() (User, bool) {
	if cu, err := GetCurUserD(); err != nil {
		return cu, false
	} else {
		return cu, true
	}
}
func UserLogin(username string, password string) bool {
	user := QueryUser(func(u *User) bool {
		if u.Username == username && u.Password == password {

			return true
		}

		return false
	})
	if len(user) == 0 {

		fmt.Println("Login: User not Exist")
		return false
	}
	SetCurUser(&user[0])
	if err := Sync(); err != nil {
		fmt.Println("Login: error occurred when set curuser")
		return false
	}
	return true
}

func UserRegister(username string, password string, email string, phone string) (bool, error) {
	user := QueryUser(func(u *User) bool {
		return u.Username == username
	})
	if len(user) == 1 {
		fmt.Println("User Register: Already exist username")
		return false, nil
	}
	CreateUser(&User{username, password, email, phone})
	if err := Sync(); err != nil {
		return true, err
	}
	return true, nil
}

func DeleteUser(username string) bool {
	DeleteUserD(func(u *User) bool {
		return u.Username == username
	})
	UpdateMeeting(
		func(m *Meeting) bool {
			return m.IsParticipator(username)
		},
		func(m *Meeting) {
			m.DeleteParticipator(username)
		})
	DeleteMeetingD(func(m *Meeting) bool {
		return m.Sponsor == username || len(m.GetParticipator()) == 0
	})
	if err := Sync(); err != nil {
		return false
	}
	return UserLogout()
}

func ListAllUser() []User {
	return QueryUser(func(u *User) bool {
		return true
	})
}

func CreateMeeting(username string, title string, startDate string, endDate string, participator []string) bool {
	for _, i := range participator {
		if username == i {
			fmt.Println("Create Meeting: sponsor can't be participator")
			return false
		}
		l := QueryUser(func(u *User) bool {
			return u.Username == i
		})
		if len(l) == 0 {
			fmt.Println("Create Meeting: no such a user : ", i)
			return false
		}
		dc := 0
		for _, j := range participator {
			if j == i {
				dc++
				if dc == 2 {
					fmt.Println("Create Meeting: duplicate participator")
					return false
				}
			}
		}
	}
	sTime, err := StringToDate(startDate)
	if err != nil {
		fmt.Println("Create Meeting: Wrong Date")
		return false
	}
	eTime, err := StringToDate(endDate)
	if err != nil {
		fmt.Println("Create Meeting: Wrong Date")
		return false
	}
	if eTime.LessThan(sTime) == true {
		fmt.Println("Create Meeting: Start Time greater than end time")
		return false
	}
	for _, p := range participator {
		l := QueryMeetingD(func(m *Meeting) bool {
			if m.Sponsor == p || m.IsParticipator(p) {
				if m.StartDate.LessOrEqual(sTime) && m.EndDate.MoreThan(sTime) {
					return true
				}
				if m.StartDate.LessThan(eTime) && m.EndDate.GreateOrEqual(eTime) {
					return true
				}
				if m.StartDate.GreateOrEqual(sTime) && m.EndDate.LessOrEqual(eTime) {
					return true
				}
			}
			return false
		})
		if len(l) > 0 {
			fmt.Println("Create Meeting: ", p, " time conflict")
			return false
		}
	}
	tu := QueryUser(func(u *User) bool {
		return u.Username == username
	})
	if len(tu) == 0 {
		fmt.Println("Create Meeting: Sponsor ", username, " not exist")
		return false
	}
	l := QueryMeetingD(func(m *Meeting) bool {
		if m.Sponsor == username || m.IsParticipator(username) {
			if m.StartDate.LessOrEqual(sTime) && m.EndDate.MoreThan(sTime) {
				return true
			}
			if m.StartDate.LessThan(eTime) && m.EndDate.GreateOrEqual(eTime) {
				return true
			}
			if m.StartDate.GreateOrEqual(sTime) && m.EndDate.LessOrEqual(eTime) {
				return true
			}
		}
		return false
	})

	if len(l) > 0 {
		fmt.Println("Create Meeting: ", username, " time conflict")
		return false
	}
	CreateMeetingD(&Meeting{username, participator, sTime, eTime, title})
	if err := Sync(); err != nil {
		return false
	}
	return true
}

func QueryMeeting(username, startDate, endDate string) ([]Meeting, bool) {
	sTime, err := StringToDate(startDate)
	var m []Meeting
	if err != nil {
		fmt.Println("Query Meeting: Wrong StartDate")
		return m, false
	}
	eTime, err := StringToDate(endDate)
	if err != nil {
		fmt.Println("Query Meeting: Wrong EndDate")
		return m, false
	}
	if eTime.LessThan(sTime) == true {
		fmt.Println("Query Meeting: Start Time greater than end time")
		return m, false
	}

	tm := QueryMeetingD(func(m *Meeting) bool {
		if m.Sponsor == username || m.IsParticipator(username) {
			if m.StartDate.LessOrEqual(sTime) && m.EndDate.MoreThan(sTime) {
				return true
			}
			if m.StartDate.LessOrEqual(eTime) && m.EndDate.GreateOrEqual(eTime) {
				return true
			}
			if m.StartDate.GreateOrEqual(sTime) && m.EndDate.LessOrEqual(eTime) {
				return true
			}
		}
		return false
	})
	return tm, true
}

func DeleteMeeting(username, title string) int {
	return DeleteMeetingD(func(m *Meeting) bool {
		return m.Sponsor == username && m.Title == title
	})
}

func QuitMeeting(username string, title string) bool {
	flag := QueryMeetingD(func(m *Meeting) bool {
		return m.Title == title && m.IsParticipator(username) == true
	})
	if len(flag) == 0 {
		return false
	}
	UpdateMeeting(func(m *Meeting) bool {
		return m.IsParticipator(username) == true && m.Title == title
	}, func(m *Meeting) {
		m.DeleteParticipator(username)
	})
	DeleteMeetingD(func(m *Meeting) bool {
		return len(m.GetParticipator()) == 0
	})
	return true
}

func ClearMeeting(username string) (int, bool) {
	cm := DeleteMeetingD(func(m *Meeting) bool {
		return m.Sponsor == username
	})
	if err := Sync(); err != nil {
		fmt.Println("Clear Meeting: Delete failed")
		return cm, false
	} else {
		return cm, true
	}
}

func AddMeetingParticipator(username string, title string, participators []string) bool {
	for _, p := range participators {
		uc := QueryUser(func(u *User) bool {
			return u.Username == p
		})
		if len(uc) == 0 {
			fmt.Println("Add Meeting Participator: No such a user: ", p)
			return false
		}
		qm := QueryMeetingD(func(m *Meeting) bool {
			return m.Sponsor == username && m.Title == title && m.IsParticipator(p)
		})
		if len(qm) != 0 {
			fmt.Println("Add Meeting Participator: ", p, "Already in meeting")
			return false
		}
	}
	mt := UpdateMeeting(func(m *Meeting) bool {
		return m.Sponsor == username && m.Title == title
	}, func(m *Meeting) {
		for _, p := range participators {
			m.AddParticipator(p)
		}
	})
	if mt == 0 {
		fmt.Println("Add Meeting Participator: no such meeting")
		return false
	}
	if err := Sync(); err != nil {
		return false
	}
	return true
}

func RemoveMeetingParticipator(username string, title string, participators []string) bool {
	for _, p := range participators {
		uc := QueryUser(func(u *User) bool {
			return u.Username == p
		})
		if len(uc) == 0 {
			fmt.Println("Remove Meeting Participator: No such a user: ", p)
			return false
		}
		qm := QueryMeetingD(func(m *Meeting) bool {
			return m.Sponsor == username && m.Title == title && m.IsParticipator(p)
		})
		if len(qm) == 0 {
			fmt.Println("Remove Meeting Participator: Not in Meeting :", p)
			return false
		}
	}
	mt := UpdateMeeting(func(m *Meeting) bool {
		return m.Sponsor == username && m.Title == title
	}, func(m *Meeting) {
		for _, p := range participators {
			m.DeleteParticipator(p)
		}
	})
	if mt == 0 {
		fmt.Println("Remove Meeting Participator: no such a meeting: ", title)
		return false
	}
	DeleteMeetingD(func(m *Meeting) bool {
		return m.Sponsor == username && len(m.GetParticipator()) == 0
	})
	if err := Sync(); err != nil {
		return false
	}
	return true
}
