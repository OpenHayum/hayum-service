package db

const createDDL = `
-- ************************************** User

CREATE TABLE IF NOT EXISTS User
(
	Id           INT PRIMARY KEY NOT NULL AUTO_INCREMENT ,
	FirstName    VARCHAR(45) NOT NULL ,
	LastName     VARCHAR(45) NOT NULL ,
	Mobile       VARCHAR(10) NOT NULL ,
	Password     VARCHAR(256) NOT NULL ,
	Email        VARCHAR(45) NOT NULL ,
	IsVerified   TINYINT NOT NULL DEFAULT 0,
	IsDeleted    TINYINT NOT NULL DEFAULT 0,
	OtpExpiresAt DATETIME,
	CreatedDate  DATETIME NOT NULL ,
	ModifiedDate DATETIME,
	DeletedDate  DATETIME 
);

-- ************************************** Artist

CREATE TABLE IF NOT EXISTS Artist
(
	Id         INT NOT NULL ,
	UserId     INT NOT NULL ,
	IsVerified TINYINT NOT NULL ,
	PRIMARY KEY (Id, UserId),
	CONSTRAINT FK_ARTIST_USER_USERID FOREIGN KEY (UserId) REFERENCES User (Id)
);

-- ************************************** Account

CREATE TABLE IF NOT EXISTS Account
(
	Id                INT NOT NULL ,
	UserId            INT NOT NULL ,
	IsPremium         TINYINT NOT NULL ,
	ViewsNum          INT NOT NULL ,
	FavoritesNum      INT NOT NULL ,
	TracksUploadedNum INT NOT NULL ,
	AccountImageLink  VARCHAR(45) NOT NULL ,
	CoverImageLink    VARCHAR(45) NOT NULL ,
	IsDeleted         TINYINT NOT NULL ,
	CreatedDate       DATETIME NOT NULL DEFAULT NOW(),
	ModifiedDate      DATETIME,
	DeletedDate       DATETIME,
	PRIMARY KEY (Id, UserId),
	CONSTRAINT FK_ACCOUNT_USER_USERID FOREIGN KEY  (UserId) REFERENCES User (Id)
);

-- ************************************** Followers

CREATE TABLE IF NOT EXISTS Follower
(
	ArtistId     INTEGER NOT NULL ,
	ArtistUserId INTEGER NOT NULL ,
	FollowerId   INTEGER NOT NULL ,
	PRIMARY KEY (FollowerId, ArtistId, ArtistUserId),
	CONSTRAINT FK_FOLLOWER_ARTIST_ID_USERID FOREIGN KEY (ArtistId, ArtistUserId) REFERENCES Artist (Id, UserId),
	CONSTRAINT FK_FOLLOWER_USER_USERID FOREIGN KEY (FollowerId) REFERENCES User (Id)
);

-- ************************************** S3Document

CREATE TABLE IF NOT EXISTS S3Document
(
	Id               INT PRIMARY KEY NOT NULL ,
	OriginalFileName VARCHAR(250) NOT NULL ,
	S3Key              VARCHAR(250) NOT NULL ,
	Bucket           VARCHAR(250) NOT NULL ,
	IsDeleted        TINYINT NOT NULL ,
	CreatedDate      DATETIME NOT NULL ,
	ModifiedDate     DATETIME,
	DeletedDate      DATETIME
);

-- ************************************** Session

CREATE TABLE IF NOT EXISTS Session
(
	SessionId VARCHAR(100) NOT NULL ,
	UserId    INT NOT NULL ,
	ExpiresAt TIMESTAMP,
	CreatedAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
	UpdatedAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
	PRIMARY KEY (SessionId, UserId),
	CONSTRAINT FK_SESSION_USER_USERID FOREIGN KEY (UserId) REFERENCES User (Id)
);
`
