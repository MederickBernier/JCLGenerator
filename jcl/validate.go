package jcl

import (
	"errors"
	"regexp"
)

// ValidateJCL ensures the JCL parameters follow correct formatting
func ValidateJCL(params JCLParameters) error {
	if err := validateJobParams(params); err != nil {
		return err
	}
	if err := validateExecParams(params); err != nil {
		return err
	}
	if err := validateDDParams(params); err != nil {
		return err
	}
	if err := validateJESParams(params); err != nil {
		return err
	}
	if params.Steps != nil {
		for _, step := range params.Steps {
			if err := ValidateStep(step); err != nil {
				return err
			}
		}
	}
	return nil
}

func validateJobParams(params JCLParameters) error {
	jobNamePattern := `^[A-Z0-9]{1,8}$`
	classPattern := `^[A-Z]{1}$`
	prtyPattern := `^[0-9]{1,2}$`
	timePattern := `^[0-9]+$`

	if params.JobName != nil && !regexp.MustCompile(jobNamePattern).MatchString(*params.JobName) {
		return errors.New("invalid Job Name (must be 1-8 uppercase letters/numbers)")
	}
	if params.Class != nil && !regexp.MustCompile(classPattern).MatchString(*params.Class) {
		return errors.New("invalid Class (must be a single uppercase letter A-Z)")
	}
	if params.Prty != nil && !regexp.MustCompile(prtyPattern).MatchString(*params.Prty) {
		return errors.New("invalid Priority (must be 1-2 digits)")
	}
	if params.Time != nil && !regexp.MustCompile(timePattern).MatchString(*params.Time) {
		return errors.New("invalid Time (must be numeric, e.g., 1440)")
	}
	return nil
}

func validateExecParams(params JCLParameters) error {
	programPattern := `^[A-Z0-9]+$`

	if params.ExecPGM != nil && !regexp.MustCompile(programPattern).MatchString(*params.ExecPGM) {
		return errors.New("invalid EXEC Program Name")
	}
	if params.ExecPROC != nil && !regexp.MustCompile(programPattern).MatchString(*params.ExecPROC) {
		return errors.New("invalid EXEC Procedure Name")
	}
	return nil
}

func validateDDParams(params JCLParameters) error {
	dsNamePattern := `^[A-Z0-9]+(\.[A-Z0-9]+)*$`
	volumePattern := `^[A-Z0-9]{1,6}$`
	dispPattern := `^(NEW|OLD|SHR|MOD)$`
	spacePattern := `^[0-9]+$`

	if params.DSName != nil && !regexp.MustCompile(dsNamePattern).MatchString(*params.DSName) {
		return errors.New("invalid Dataset Name (DSN)")
	}
	if params.Volume != nil && !regexp.MustCompile(volumePattern).MatchString(*params.Volume) {
		return errors.New("invalid Volume Serial")
	}
	if params.DISP != nil && !regexp.MustCompile(dispPattern).MatchString(*params.DISP) {
		return errors.New("invalid DISP value (must be NEW, OLD, SHR, or MOD)")
	}
	if params.Space != nil && !regexp.MustCompile(spacePattern).MatchString(*params.Space) {
		return errors.New("invalid Space Allocation (must be numeric)")
	}
	return nil
}

func validateJESParams(params JCLParameters) error {
	jobParmPattern := `^[A-Z0-9]+$`

	if params.JobParm != nil && !regexp.MustCompile(jobParmPattern).MatchString(*params.JobParm) {
		return errors.New("invalid JOBPARM value")
	}
	return nil
}

func ValidateStep(step JCLStep) error {
	stepNamePattern := `^[A-Z0-9]{1,8}$`

	if step.StepName != nil && !regexp.MustCompile(stepNamePattern).MatchString(*step.StepName) {
		return errors.New("invalid Step Name (must be 1-8 uppercase letters/numbers)")
	}
	return nil
}
