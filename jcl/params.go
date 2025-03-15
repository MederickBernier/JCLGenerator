package jcl

// JCLParameters defines the structure for JCL parameterization.
type JCLParameters struct {
	// JOB Statement Parameters
	JobName  *string // Job name (1-8 characters)
	Class    *string // Execution class (A-Z)
	MsgClass *string // Message class for output
	MsgLevel *string // Controls job log messages
	Notify   *string // User to notify when the job completes
	Prty     *string // Job priority (0-15)
	Region   *string // Memory allocation
	Restart  *string // Restart instructions
	Time     *string // CPU time limit
	TypeRun  *string // Special job execution (HOLD, SCAN)
	User     *string // User ID submitting the job
	Password *string // User password (if required)

	// EXEC Statement Parameters (Per Step)
	ExecPGM  *string // Program to execute
	ExecPROC *string // Procedure to execute
	ExecParm *string // Parameters passed to the program
	Cond     *string // Condition codes for step execution
	AddrSpc  *string // Address space requirements
	TimeStep *string // Time limit for the step

	// DD Statement Parameters (Dataset Handling)
	DSName    *string // Dataset name
	DISP      *string // Disposition (NEW, OLD, SHR, MOD)
	UNIT      *string // Device type
	Volume    *string // Volume serial number
	Space     *string // Space allocation
	SpaceUnit *string // Space unit (CYL, TRK, BYTES)
	DCB       *string // Data control block attributes
	LRECL     *string // Logical record length
	BLKSIZE   *string // Block size
	RecFM     *string // Record format (FB, VB, etc.)
	ExpDT     *string // Expiration date (YYYY/MM/DD)
	RetPD     *string // Retention period (days)
	SysOut    *string // SYSOUT class (A-Z)
	SysIn     *string // SYSIN input
	Dummy     *string // DUMMY dataset indicator

	// SYSIN/SYSOUT Statement Parameters
	SysInDD  *string // SYSIN DD statement for input
	SysOutDD *string // SYSOUT DD statement for output

	// JES2/JES3 Control Statements
	JobParm *string // JOBPARM for JES2
	Output  *string // Output processing options
	Route   *string // Route job/output to a system
	Main    *string // MAIN statement for JES3
	Format  *string // FORMAT statement for JES3
	DataSet *string // In-stream dataset for JES3

	// Miscellaneous Parameters
	AddrSpcExec *string // Address space for execution
	Bytes       *string // Max bytes for job
	Lines       *string // Max printed lines
	Pages       *string // Max printed pages
	Perform     *string // Performance group/service class
	RD          *string // Restart disposition
	Chars       *string // Printer character arrangement
	FCB         *string // Forms Control Buffer
	Flash       *string // Form overlays
	Forms       *string // Forms name for printing
	Modify      *string // Modifications to dataset attributes
	PrtAttrs    *string // Printer attributes
	UCS         *string // Universal Character Set for printing
	Writer      *string // External writer for output processing

	// Job Steps (Future Enhancement in v0.3)
	Steps []JCLStep // Holds multiple execution steps

	// System-Level Overrides
	SystemOverrides *JCLSystemOverrides // Advanced tuning
}

// JCLStep represents an execution step inside a JCL job.
type JCLStep struct {
	StepName *string // Name of the step
	ExecPGM  *string // Program to execute
	ExecPROC *string // Procedure to execute
	ExecParm *string // Parameters for execution
	Cond     *string // Execution conditions
	Time     *string // Time limit
}

// JCLSystemOverrides provides system-level settings for the JCL execution.
type JCLSystemOverrides struct {
	AddrSpace *string // Address space allocation
	Perform   *string // Performance tuning settings
	RD        *string // Restart disposition (system-wide)
}
