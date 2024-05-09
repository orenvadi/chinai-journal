package grpcauth

import (
	"context"

	"github.com/orenvadi/auth-grpc/protos/gen/go/proto/sso"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (s *serverAPI) GetAttendanceJournal(
	ctx context.Context,
	req *sso.GetAttendanceJournalRequest,
) (
	res *sso.GetAttendanceJournalResponse,
	err error,
) {
	journal, err := s.confCodes.GetAttendanceJournal(ctx, req.Date.AsTime())
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	journalLessons := []*sso.GetAttendanceJournalResponse_AttendanceJournalLesson{}

	for _, lessons := range journal {

		journalLines := []*sso.GetAttendanceJournalResponse_AttendanceJournalLesson_AttendanceJournalLine{}

		for _, line := range lessons.AttendanceJournalLine {
			journalLines = append(journalLines, &sso.GetAttendanceJournalResponse_AttendanceJournalLesson_AttendanceJournalLine{
				Number:      int32(line.Number),
				StudentName: line.StudentName,
				IsAttended:  line.IsAttended,
			})
		}

		journalLessons = append(journalLessons, &sso.GetAttendanceJournalResponse_AttendanceJournalLesson{
			TimeSlot:               timestamppb.New(lessons.TimeSlot),
			Lesson:                 lessons.Lesson,
			Group:                  lessons.Group,
			AttendanceJournalLines: journalLines,
		})
	}

	return &sso.GetAttendanceJournalResponse{AttendanceJournalLessons: journalLessons}, nil
}

func (s *serverAPI) GetAttendanceLessons(
	ctx context.Context,
	req *sso.GetAttendanceLessonsRequest) (
	res *sso.GetAttendanceLessonsResponse,
	err error,
) {
	lessons, err := s.confCodes.GetAttendanceLessons(ctx, req.Date.AsTime())
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	lessonsLines := []*sso.GetAttendanceLessonsResponse_AttendanceLessonsLine{}

	for _, lesson := range lessons.AttendanceLessons {
		lessonsLines = append(lessonsLines, &sso.GetAttendanceLessonsResponse_AttendanceLessonsLine{
			TimeSlot:   timestamppb.New(lesson.TimeSlot),
			Group:      lesson.Group,
			Subject:    lesson.Subject,
			IsAttended: lesson.IsAttended,
		})
	}

	res = &sso.GetAttendanceLessonsResponse{
		Date:                  timestamppb.New(lessons.Date),
		AttendanceLessonsLine: lessonsLines,
	}

	return res, nil
}
