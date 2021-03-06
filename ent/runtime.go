// Code generated by entc, DO NOT EDIT.

package ent

import (
	"street/ent/account"
	"street/ent/comment"
	"street/ent/episode"
	"street/ent/file"
	"street/ent/profile"
	"street/ent/schema"
	"street/ent/series"
	"street/ent/token"
	"time"

	"github.com/google/uuid"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	accountMixin := schema.Account{}.Mixin()
	accountMixinFields0 := accountMixin[0].Fields()
	_ = accountMixinFields0
	accountMixinFields1 := accountMixin[1].Fields()
	_ = accountMixinFields1
	accountFields := schema.Account{}.Fields()
	_ = accountFields
	// accountDescSid is the schema descriptor for sid field.
	accountDescSid := accountMixinFields0[0].Descriptor()
	// account.DefaultSid holds the default value on creation for the sid field.
	account.DefaultSid = accountDescSid.Default.(func() schema.ID)
	// accountDescCreateTime is the schema descriptor for create_time field.
	accountDescCreateTime := accountMixinFields1[0].Descriptor()
	// account.DefaultCreateTime holds the default value on creation for the create_time field.
	account.DefaultCreateTime = accountDescCreateTime.Default.(func() time.Time)
	// accountDescUpdateTime is the schema descriptor for update_time field.
	accountDescUpdateTime := accountMixinFields1[1].Descriptor()
	// account.DefaultUpdateTime holds the default value on creation for the update_time field.
	account.DefaultUpdateTime = accountDescUpdateTime.Default.(func() time.Time)
	// account.UpdateDefaultUpdateTime holds the default value on update for the update_time field.
	account.UpdateDefaultUpdateTime = accountDescUpdateTime.UpdateDefault.(func() time.Time)
	// accountDescEmail is the schema descriptor for email field.
	accountDescEmail := accountFields[0].Descriptor()
	// account.EmailValidator is a validator for the "email" field. It is called by the builders before save.
	account.EmailValidator = func() func(string) error {
		validators := accountDescEmail.Validators
		fns := [...]func(string) error{
			validators[0].(func(string) error),
			validators[1].(func(string) error),
		}
		return func(email string) error {
			for _, fn := range fns {
				if err := fn(email); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	// accountDescID is the schema descriptor for id field.
	accountDescID := accountMixinFields0[1].Descriptor()
	// account.DefaultID holds the default value on creation for the id field.
	account.DefaultID = accountDescID.Default.(func() uuid.UUID)
	commentMixin := schema.Comment{}.Mixin()
	commentMixinFields0 := commentMixin[0].Fields()
	_ = commentMixinFields0
	commentMixinFields1 := commentMixin[1].Fields()
	_ = commentMixinFields1
	commentFields := schema.Comment{}.Fields()
	_ = commentFields
	// commentDescSid is the schema descriptor for sid field.
	commentDescSid := commentMixinFields0[0].Descriptor()
	// comment.DefaultSid holds the default value on creation for the sid field.
	comment.DefaultSid = commentDescSid.Default.(func() schema.ID)
	// commentDescCreateTime is the schema descriptor for create_time field.
	commentDescCreateTime := commentMixinFields1[0].Descriptor()
	// comment.DefaultCreateTime holds the default value on creation for the create_time field.
	comment.DefaultCreateTime = commentDescCreateTime.Default.(func() time.Time)
	// commentDescUpdateTime is the schema descriptor for update_time field.
	commentDescUpdateTime := commentMixinFields1[1].Descriptor()
	// comment.DefaultUpdateTime holds the default value on creation for the update_time field.
	comment.DefaultUpdateTime = commentDescUpdateTime.Default.(func() time.Time)
	// comment.UpdateDefaultUpdateTime holds the default value on update for the update_time field.
	comment.UpdateDefaultUpdateTime = commentDescUpdateTime.UpdateDefault.(func() time.Time)
	// commentDescContent is the schema descriptor for content field.
	commentDescContent := commentFields[0].Descriptor()
	// comment.ContentValidator is a validator for the "content" field. It is called by the builders before save.
	comment.ContentValidator = commentDescContent.Validators[0].(func(string) error)
	// commentDescPath is the schema descriptor for path field.
	commentDescPath := commentFields[1].Descriptor()
	// comment.DefaultPath holds the default value on creation for the path field.
	comment.DefaultPath = commentDescPath.Default.(string)
	// commentDescID is the schema descriptor for id field.
	commentDescID := commentMixinFields0[1].Descriptor()
	// comment.DefaultID holds the default value on creation for the id field.
	comment.DefaultID = commentDescID.Default.(func() uuid.UUID)
	episodeMixin := schema.Episode{}.Mixin()
	episodeMixinFields0 := episodeMixin[0].Fields()
	_ = episodeMixinFields0
	episodeMixinFields1 := episodeMixin[1].Fields()
	_ = episodeMixinFields1
	episodeFields := schema.Episode{}.Fields()
	_ = episodeFields
	// episodeDescSid is the schema descriptor for sid field.
	episodeDescSid := episodeMixinFields0[0].Descriptor()
	// episode.DefaultSid holds the default value on creation for the sid field.
	episode.DefaultSid = episodeDescSid.Default.(func() schema.ID)
	// episodeDescCreateTime is the schema descriptor for create_time field.
	episodeDescCreateTime := episodeMixinFields1[0].Descriptor()
	// episode.DefaultCreateTime holds the default value on creation for the create_time field.
	episode.DefaultCreateTime = episodeDescCreateTime.Default.(func() time.Time)
	// episodeDescUpdateTime is the schema descriptor for update_time field.
	episodeDescUpdateTime := episodeMixinFields1[1].Descriptor()
	// episode.DefaultUpdateTime holds the default value on creation for the update_time field.
	episode.DefaultUpdateTime = episodeDescUpdateTime.Default.(func() time.Time)
	// episode.UpdateDefaultUpdateTime holds the default value on update for the update_time field.
	episode.UpdateDefaultUpdateTime = episodeDescUpdateTime.UpdateDefault.(func() time.Time)
	// episodeDescCover is the schema descriptor for cover field.
	episodeDescCover := episodeFields[0].Descriptor()
	// episode.CoverValidator is a validator for the "cover" field. It is called by the builders before save.
	episode.CoverValidator = episodeDescCover.Validators[0].(func(string) error)
	// episodeDescTitle is the schema descriptor for title field.
	episodeDescTitle := episodeFields[1].Descriptor()
	// episode.TitleValidator is a validator for the "title" field. It is called by the builders before save.
	episode.TitleValidator = episodeDescTitle.Validators[0].(func(string) error)
	// episodeDescContent is the schema descriptor for content field.
	episodeDescContent := episodeFields[2].Descriptor()
	// episode.ContentValidator is a validator for the "content" field. It is called by the builders before save.
	episode.ContentValidator = episodeDescContent.Validators[0].(func(string) error)
	// episodeDescFiles is the schema descriptor for files field.
	episodeDescFiles := episodeFields[3].Descriptor()
	// episode.DefaultFiles holds the default value on creation for the files field.
	episode.DefaultFiles = episodeDescFiles.Default.(func() schema.Medias)
	// episodeDescID is the schema descriptor for id field.
	episodeDescID := episodeMixinFields0[1].Descriptor()
	// episode.DefaultID holds the default value on creation for the id field.
	episode.DefaultID = episodeDescID.Default.(func() uuid.UUID)
	fileMixin := schema.File{}.Mixin()
	fileMixinFields0 := fileMixin[0].Fields()
	_ = fileMixinFields0
	fileMixinFields1 := fileMixin[1].Fields()
	_ = fileMixinFields1
	fileFields := schema.File{}.Fields()
	_ = fileFields
	// fileDescSid is the schema descriptor for sid field.
	fileDescSid := fileMixinFields0[0].Descriptor()
	// file.DefaultSid holds the default value on creation for the sid field.
	file.DefaultSid = fileDescSid.Default.(func() schema.ID)
	// fileDescCreateTime is the schema descriptor for create_time field.
	fileDescCreateTime := fileMixinFields1[0].Descriptor()
	// file.DefaultCreateTime holds the default value on creation for the create_time field.
	file.DefaultCreateTime = fileDescCreateTime.Default.(func() time.Time)
	// fileDescUpdateTime is the schema descriptor for update_time field.
	fileDescUpdateTime := fileMixinFields1[1].Descriptor()
	// file.DefaultUpdateTime holds the default value on creation for the update_time field.
	file.DefaultUpdateTime = fileDescUpdateTime.Default.(func() time.Time)
	// file.UpdateDefaultUpdateTime holds the default value on update for the update_time field.
	file.UpdateDefaultUpdateTime = fileDescUpdateTime.UpdateDefault.(func() time.Time)
	// fileDescFilename is the schema descriptor for filename field.
	fileDescFilename := fileFields[0].Descriptor()
	// file.DefaultFilename holds the default value on creation for the filename field.
	file.DefaultFilename = fileDescFilename.Default.(string)
	// file.FilenameValidator is a validator for the "filename" field. It is called by the builders before save.
	file.FilenameValidator = func() func(string) error {
		validators := fileDescFilename.Validators
		fns := [...]func(string) error{
			validators[0].(func(string) error),
			validators[1].(func(string) error),
		}
		return func(filename string) error {
			for _, fn := range fns {
				if err := fn(filename); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	// fileDescPath is the schema descriptor for path field.
	fileDescPath := fileFields[1].Descriptor()
	// file.DefaultPath holds the default value on creation for the path field.
	file.DefaultPath = fileDescPath.Default.(string)
	// file.PathValidator is a validator for the "path" field. It is called by the builders before save.
	file.PathValidator = func() func(string) error {
		validators := fileDescPath.Validators
		fns := [...]func(string) error{
			validators[0].(func(string) error),
			validators[1].(func(string) error),
		}
		return func(_path string) error {
			for _, fn := range fns {
				if err := fn(_path); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	// fileDescMime is the schema descriptor for mime field.
	fileDescMime := fileFields[2].Descriptor()
	// file.DefaultMime holds the default value on creation for the mime field.
	file.DefaultMime = fileDescMime.Default.(string)
	// file.MimeValidator is a validator for the "mime" field. It is called by the builders before save.
	file.MimeValidator = func() func(string) error {
		validators := fileDescMime.Validators
		fns := [...]func(string) error{
			validators[0].(func(string) error),
			validators[1].(func(string) error),
		}
		return func(mime string) error {
			for _, fn := range fns {
				if err := fn(mime); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	// fileDescStatus is the schema descriptor for status field.
	fileDescStatus := fileFields[4].Descriptor()
	// file.DefaultStatus holds the default value on creation for the status field.
	file.DefaultStatus = fileDescStatus.Default.(string)
	// file.StatusValidator is a validator for the "status" field. It is called by the builders before save.
	file.StatusValidator = fileDescStatus.Validators[0].(func(string) error)
	// fileDescNote is the schema descriptor for note field.
	fileDescNote := fileFields[5].Descriptor()
	// file.NoteValidator is a validator for the "note" field. It is called by the builders before save.
	file.NoteValidator = fileDescNote.Validators[0].(func(string) error)
	// fileDescID is the schema descriptor for id field.
	fileDescID := fileMixinFields0[1].Descriptor()
	// file.DefaultID holds the default value on creation for the id field.
	file.DefaultID = fileDescID.Default.(func() uuid.UUID)
	profileMixin := schema.Profile{}.Mixin()
	profileMixinFields0 := profileMixin[0].Fields()
	_ = profileMixinFields0
	profileMixinFields1 := profileMixin[1].Fields()
	_ = profileMixinFields1
	profileFields := schema.Profile{}.Fields()
	_ = profileFields
	// profileDescSid is the schema descriptor for sid field.
	profileDescSid := profileMixinFields0[0].Descriptor()
	// profile.DefaultSid holds the default value on creation for the sid field.
	profile.DefaultSid = profileDescSid.Default.(func() schema.ID)
	// profileDescCreateTime is the schema descriptor for create_time field.
	profileDescCreateTime := profileMixinFields1[0].Descriptor()
	// profile.DefaultCreateTime holds the default value on creation for the create_time field.
	profile.DefaultCreateTime = profileDescCreateTime.Default.(func() time.Time)
	// profileDescUpdateTime is the schema descriptor for update_time field.
	profileDescUpdateTime := profileMixinFields1[1].Descriptor()
	// profile.DefaultUpdateTime holds the default value on creation for the update_time field.
	profile.DefaultUpdateTime = profileDescUpdateTime.Default.(func() time.Time)
	// profile.UpdateDefaultUpdateTime holds the default value on update for the update_time field.
	profile.UpdateDefaultUpdateTime = profileDescUpdateTime.UpdateDefault.(func() time.Time)
	// profileDescTitle is the schema descriptor for title field.
	profileDescTitle := profileFields[0].Descriptor()
	// profile.TitleValidator is a validator for the "title" field. It is called by the builders before save.
	profile.TitleValidator = func() func(string) error {
		validators := profileDescTitle.Validators
		fns := [...]func(string) error{
			validators[0].(func(string) error),
			validators[1].(func(string) error),
		}
		return func(title string) error {
			for _, fn := range fns {
				if err := fn(title); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	// profileDescCall is the schema descriptor for call field.
	profileDescCall := profileFields[1].Descriptor()
	// profile.CallValidator is a validator for the "call" field. It is called by the builders before save.
	profile.CallValidator = func() func(string) error {
		validators := profileDescCall.Validators
		fns := [...]func(string) error{
			validators[0].(func(string) error),
			validators[1].(func(string) error),
		}
		return func(call string) error {
			for _, fn := range fns {
				if err := fn(call); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	// profileDescCategory is the schema descriptor for category field.
	profileDescCategory := profileFields[2].Descriptor()
	// profile.CategoryValidator is a validator for the "category" field. It is called by the builders before save.
	profile.CategoryValidator = func() func(string) error {
		validators := profileDescCategory.Validators
		fns := [...]func(string) error{
			validators[0].(func(string) error),
			validators[1].(func(string) error),
		}
		return func(category string) error {
			for _, fn := range fns {
				if err := fn(category); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	// profileDescAvatar is the schema descriptor for avatar field.
	profileDescAvatar := profileFields[3].Descriptor()
	// profile.AvatarValidator is a validator for the "avatar" field. It is called by the builders before save.
	profile.AvatarValidator = profileDescAvatar.Validators[0].(func(string) error)
	// profileDescID is the schema descriptor for id field.
	profileDescID := profileMixinFields0[1].Descriptor()
	// profile.DefaultID holds the default value on creation for the id field.
	profile.DefaultID = profileDescID.Default.(func() uuid.UUID)
	seriesMixin := schema.Series{}.Mixin()
	seriesMixinFields0 := seriesMixin[0].Fields()
	_ = seriesMixinFields0
	seriesMixinFields1 := seriesMixin[1].Fields()
	_ = seriesMixinFields1
	seriesFields := schema.Series{}.Fields()
	_ = seriesFields
	// seriesDescSid is the schema descriptor for sid field.
	seriesDescSid := seriesMixinFields0[0].Descriptor()
	// series.DefaultSid holds the default value on creation for the sid field.
	series.DefaultSid = seriesDescSid.Default.(func() schema.ID)
	// seriesDescCreateTime is the schema descriptor for create_time field.
	seriesDescCreateTime := seriesMixinFields1[0].Descriptor()
	// series.DefaultCreateTime holds the default value on creation for the create_time field.
	series.DefaultCreateTime = seriesDescCreateTime.Default.(func() time.Time)
	// seriesDescUpdateTime is the schema descriptor for update_time field.
	seriesDescUpdateTime := seriesMixinFields1[1].Descriptor()
	// series.DefaultUpdateTime holds the default value on creation for the update_time field.
	series.DefaultUpdateTime = seriesDescUpdateTime.Default.(func() time.Time)
	// series.UpdateDefaultUpdateTime holds the default value on update for the update_time field.
	series.UpdateDefaultUpdateTime = seriesDescUpdateTime.UpdateDefault.(func() time.Time)
	// seriesDescTitle is the schema descriptor for title field.
	seriesDescTitle := seriesFields[0].Descriptor()
	// series.TitleValidator is a validator for the "title" field. It is called by the builders before save.
	series.TitleValidator = func() func(string) error {
		validators := seriesDescTitle.Validators
		fns := [...]func(string) error{
			validators[0].(func(string) error),
			validators[1].(func(string) error),
		}
		return func(title string) error {
			for _, fn := range fns {
				if err := fn(title); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	// seriesDescType is the schema descriptor for type field.
	seriesDescType := seriesFields[1].Descriptor()
	// series.TypeValidator is a validator for the "type" field. It is called by the builders before save.
	series.TypeValidator = seriesDescType.Validators[0].(func(string) error)
	// seriesDescID is the schema descriptor for id field.
	seriesDescID := seriesMixinFields0[1].Descriptor()
	// series.DefaultID holds the default value on creation for the id field.
	series.DefaultID = seriesDescID.Default.(func() uuid.UUID)
	tokenMixin := schema.Token{}.Mixin()
	tokenMixinFields0 := tokenMixin[0].Fields()
	_ = tokenMixinFields0
	tokenMixinFields1 := tokenMixin[1].Fields()
	_ = tokenMixinFields1
	tokenFields := schema.Token{}.Fields()
	_ = tokenFields
	// tokenDescSid is the schema descriptor for sid field.
	tokenDescSid := tokenMixinFields0[0].Descriptor()
	// token.DefaultSid holds the default value on creation for the sid field.
	token.DefaultSid = tokenDescSid.Default.(func() schema.ID)
	// tokenDescCreateTime is the schema descriptor for create_time field.
	tokenDescCreateTime := tokenMixinFields1[0].Descriptor()
	// token.DefaultCreateTime holds the default value on creation for the create_time field.
	token.DefaultCreateTime = tokenDescCreateTime.Default.(func() time.Time)
	// tokenDescUpdateTime is the schema descriptor for update_time field.
	tokenDescUpdateTime := tokenMixinFields1[1].Descriptor()
	// token.DefaultUpdateTime holds the default value on creation for the update_time field.
	token.DefaultUpdateTime = tokenDescUpdateTime.Default.(func() time.Time)
	// token.UpdateDefaultUpdateTime holds the default value on update for the update_time field.
	token.UpdateDefaultUpdateTime = tokenDescUpdateTime.UpdateDefault.(func() time.Time)
	// tokenDescBody is the schema descriptor for body field.
	tokenDescBody := tokenFields[0].Descriptor()
	// token.BodyValidator is a validator for the "body" field. It is called by the builders before save.
	token.BodyValidator = func() func(string) error {
		validators := tokenDescBody.Validators
		fns := [...]func(string) error{
			validators[0].(func(string) error),
			validators[1].(func(string) error),
		}
		return func(body string) error {
			for _, fn := range fns {
				if err := fn(body); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	// tokenDescType is the schema descriptor for type field.
	tokenDescType := tokenFields[1].Descriptor()
	// token.TypeValidator is a validator for the "type" field. It is called by the builders before save.
	token.TypeValidator = func() func(string) error {
		validators := tokenDescType.Validators
		fns := [...]func(string) error{
			validators[0].(func(string) error),
			validators[1].(func(string) error),
		}
		return func(_type string) error {
			for _, fn := range fns {
				if err := fn(_type); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	// tokenDescID is the schema descriptor for id field.
	tokenDescID := tokenMixinFields0[1].Descriptor()
	// token.DefaultID holds the default value on creation for the id field.
	token.DefaultID = tokenDescID.Default.(func() uuid.UUID)
}
