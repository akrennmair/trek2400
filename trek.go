package main

const (
	/* galactic parameters */
	NSECTS = 10 /* dimensions of quadrant in sectors */
	NQUADS = 8  /* dimension of galazy in quadrants */
	NINHAB = 32 /* number of quadrants which are inhabited */

	NEVENTS  = 12
	MAXBASES = 9

	NBANKS = 6 /* number of phaser banks */

	MAXEVENTS = 25
	MAXDISTR  = 5 /* maximum concurrent distress calls */

	Q_DISTRESSED = 0200
	Q_SYSTEM     = 077

	MAXKLQUAD = 9 /* maximum klingons per quadrant */

	GREEN  = 0
	DOCKED = 1
	YELLOW = 2
	RED    = 3

	/* device tokens */
	WARP     = 0  /* warp engines */
	SRSCAN   = 1  /* short range scanners */
	LRSCAN   = 2  /* long range scanners */
	PHASER   = 3  /* phaser control */
	TORPED   = 4  /* photon torpedo control */
	IMPULSE  = 5  /* impulse engines */
	SHIELD   = 6  /* shield control */
	COMPUTER = 7  /* on board computer */
	SSRADIO  = 8  /* subspace radio */
	LIFESUP  = 9  /* life support systems */
	SINS     = 10 /* Space Inertial Navigation System */
	CLOAK    = 11 /* cloaking device */
	XPORTER  = 12 /* transporter */
	SHUTTLE  = 13 /* shuttlecraft */

	/* Klingon move indicies */
	KM_OB = 0 /* Old quadrant, Before attack */
	KM_OA = 1 /* Old quadrant, After attack */
	KM_EB = 2 /* Enter quadrant, Before attack */
	KM_EA = 3 /* Enter quadrant, After attack */
	KM_LB = 4 /* Leave quadrant, Before attack */
	KM_LA = 5 /* Leave quadrant, After attack */

	/***************************  EVENTS  ****************************/
	E_LRTB   = 1  /* long range tractor beam */
	E_KATSB  = 2  /* Klingon attacks starbase */
	E_KDESB  = 3  /* Klingon destroys starbase */
	E_ISSUE  = 4  /* distress call is issued */
	E_ENSLV  = 5  /* Klingons enslave a quadrant */
	E_REPRO  = 6  /* a Klingon is reproduced */
	E_FIXDV  = 7  /* fix a device */
	E_ATTACK = 8  /* Klingon attack during rest period */
	E_SNAP   = 9  /* take a snapshot for time warp */
	E_SNOVA  = 10 /* supernova occurs */

	E_GHOST  = 0100 /* ghost of a distress call if ssradio out */
	E_HIDDEN = 0200 /* event that is unreportable because ssradio out */
	E_EVENT  = 077  /* mask to get event code */

	/* defines for sector map  (below) */
	EMPTY      = '.'
	STAR       = '*'
	BASE       = '#'
	ENTERPRISE = 'E'
	QUEENE     = 'Q'
	KLINGON    = 'K'
	INHABIT    = '@'
	HOLE       = ' '

	/* you lose codes */
	L_NOTIME   = 1  /* ran out of time */
	L_NOENGY   = 2  /* ran out of energy */
	L_DSTRYD   = 3  /* destroyed by a Klingon */
	L_NEGENB   = 4  /* ran into the negative energy barrier */
	L_SUICID   = 5  /* destroyed in a nova */
	L_SNOVA    = 6  /* destroyed in a supernova */
	L_NOLIFE   = 7  /* life support died (so did you) */
	L_NOHELP   = 8  /* you could not be rematerialized */
	L_TOOFAST  = 9  /* pretty stupid going at warp 10 */
	L_STAR     = 10 /* ran into a star */
	L_DSTRCT   = 11 /* self destructed */
	L_CAPTURED = 12 /* captured by Klingons */
	L_NOCREW   = 13 /* you ran out of crew */

	TOOLARGE = 1e50
)

type Game struct {
	killk     int    /* number of klingons killed */
	deaths    int    /* number of deaths onboard Enterprise */
	negenbar  int    /* number of hits on negative energy barrier */
	killb     int    /* number of starbases killed */
	kills     int    /* number of stars killed */
	skill     int    /* skill rating of player */
	length    int    /* length of game */
	killed    bool   /* set if you were killed */
	killinhab int    /* number of inhabited starsystems killed */
	tourn     bool   /* set if a tournament game */
	passwd    string /* game password */
	snap      bool   /* set if snapshot taken */
	helps     int    /* number of help calls */
	captives  int    /* total number of captives taken */
}

var game Game

type Param struct {
	bases       int             /* number of starbases */
	klings      int             /* number of klingons */
	date        float64         /* stardate */
	time        float64         /* time left */
	resource    float64         /* Federation resources */
	energy      int             /* starship's energy */
	shield      int             /* energy in shields */
	reserves    float64         /* life support reserves */
	crew        int             /* size of ship's complement */
	brigfree    int             /* max possible number of captives */
	torped      int             /* photon torpedos */
	damfac      map[int]float64 /* damage factor */
	dockfac     float64         /* docked repair time factor */
	regenfac    float64         /* regeneration factor */
	stopengy    int             /* energy to do emergency stop */
	shupengy    int             /* energy to put up shields */
	klingpwr    int             /* Klingon initial power */
	warptime    int             /* time chewer multiplier */
	phasfac     float64         /* Klingon phaser power eater factor */
	moveprob    map[int]float64 /* probability that a Klingon moves */
	movefac     map[int]float64 /* Klingon move distance multiplier */
	eventdly    map[int]float64 /* event time multipliers */
	navigcrud   []float64       /* navigation crudup factor */
	cloakenergy int             /* cloaking device energy per stardate */
	damprob     map[int]float64 /* damage probability */
	hitfac      float64         /* Klingon attack factor */
	klingcrew   int             /* number of Klingons in a crew */
	srndrprob   float64         /* surrender probability */
	energylow   int             /* low energy mark (cond YELLOW) */
}

var param Param

type Now struct {
	bases      int             /* number of starbases */
	klings     int             /* number of klingons */
	date       float64         /* stardate */
	time       float64         /* time left */
	resource   float64         /* Federation resources */
	distressed int             /* number of currently distressed quadrants */
	eventptr   [NEVENTS]*event /* pointer to event structs */
	base       [MAXBASES]xy    /* locations of starbases */
}

var now Now

type xy struct {
	x, y int /* coordinates */
}

type event struct {
	x, y       int     /* coordinates */
	date       float64 /* trap stardate */
	evcode     int     /* event type */
	systemname int     /* starsystem name */
}

var eventList [MAXEVENTS]event

type Ship struct {
	warp      float64 /* warp factor */
	warp2     float64 /* warp factor squared */
	warp3     float64 /* warp factor cubed */
	shldup    bool    /* shield up flag */
	cloaked   bool    /* set if cloaking device on */
	energy    int     /* starship's energy */
	shield    int     /* energy in shields */
	reserves  float64 /* life support reserves */
	crew      int     /* ship's complement */
	brigfree  int     /* space left in brig */
	torped    int     /* torpedoes */
	cloakgood bool    /* set if we have moved */
	quadx     int     /* quadrant x coord */
	quady     int     /* quadrant y coord */
	sectx     int     /* sector x coord */
	secty     int     /* sector y coord */
	cond      int     /* condition code */
	/* sinsbad is set if SINS is working but not calibrated */
	sinsbad    bool   /* Space Inertial Navigation System condition */
	shipname   string /* name of current starship */
	ship       byte   /* current starship */
	distressed int    /* number of distress calls */
}

var ship Ship

type device struct {
	name   string /* device name */
	person string /* the person who fixes it */
}

var devices = []device{
	{"warp drive", "Scotty"},
	{"S.R. scanners", "Scotty"},
	{"L.R. scanners", "Scotty"},
	{"phasers", "Sulu"},
	{"photon tubes", "Sulu"},
	{"impulse engines", "Scotty"},
	{"shield control", "Sulu"},
	{"computer", "Spock"},
	{"subspace radio", "Uhura"},
	{"life support", "Scotty"},
	{"navigation system", "Chekov"},
	{"cloaking device", "Scotty"},
	{"transporter", "Scotty"},
	{"shuttlecraft", "Scotty"},
}

type quadrant struct {
	bases       int /* number of bases in this quadrant */
	klings      int /* number of Klingons in this quadrant */
	holes       int /* number of black holes in this quadrant */
	scanned     int /* star chart entry (see below) */
	stars       int /* number of stars in this quadrant */
	qsystemname int /* starsystem name (see below) */
}

var quad [NQUADS][NQUADS]quadrant

type Move struct {
	free    bool    /* set if a move is free */
	endgame int     /* end of game flag */
	shldchg bool    /* set if shields changed this move */
	newquad int     /* set if just entered this quadrant */
	resting bool    /* set if this move is a rest */
	time    float64 /* time used this move */
}

var move Move

type Etc struct {
	klingon [MAXKLQUAD]kling /* sorted Klingon list */
	nkling  int              /* number of Klingons in this sector */
	/* < 0 means automatic override mode */
	starbase   xy       /* starbase in current quadrant */
	snapshot   snapshot /*snapshot for time warp */
	statreport bool     /* set to get a status report on a srscan */
}

type snapshot struct {
	quad  [NQUADS][NQUADS]quadrant
	event [MAXEVENTS]event
	now   Now
}

var etc Etc

type kling struct {
	x, y    int     /* coordinates */
	power   int     /* power left */
	dist    float64 /* distance to Enterprise */
	avgdist float64 /* average over this move */
	srndreq bool    /* set if surrender has been requested */
}

var sect [NSECTS][NSECTS]byte
