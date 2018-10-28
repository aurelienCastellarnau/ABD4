/*
 * File: iserveroption.go
 * Project: ABD4/VMD Escape Game
 * File Created: Thursday, 11th October 2018 4:28:51 pm
 * Author: billaud_j castel_a masera_m
 * Contact: (billaud_j@etna-alternance.net castel_a@etna-alternance.net masera_m@etna-alternance.net)
 * -----
 * Last Modified: Sunday, 28th October 2018 1:59:22 pm
 * Modified By: Aurélien Castellarnau
 * -----
 * Copyright © 2018 - 2018 billaud_j castel_a masera_m, ETNA - VDM EscapeGame API
 */

package context

// IServerOption :
type IServerOption interface {
	SetEnv(string)            // set environnement prod|dev|test
	SetLogpath(string)        // set path for log from exe folder
	SetDatabaseType(string)   // set kind of database mongo|bolt
	SetDatapath(string)       // for bolt database, define the .dat folder from exe folder
	SetMongoIP(ip string)     // set mongo server instance ip
	SetMongoPort(port string) // set mongo server instance port
	GetExeFolder() string     // return exe folder absolute path
	GetEnv() string           // return environnement prod|dev|test
	GetEmbedES() bool         // return a bool defining if we try to connect to elasticsearch instance
	GetEs() string            // return elasticsearch server instance address
	GetIndex() bool           // return a boolean defining if we want to create indexes
	GetReindex() bool         // return a boolean defining if we want to remove/create indexes
	GetRmindex() bool         // return a boolean defining if we want to remove indexes
	GetLogpath() string       // return the absolute path to log folder
	GetDatabaseType() string  // return the kind of database mongo|bolt
	GetDatapath() string      // return the absolute path to .dat folder in bolt database context
	GetAddress() string       // return the API server address
	GetPort() string          // return the API port
	GetIP() string            // return the API ip
	GetMongoIP() string       // return the mongo server instance ip
	GetMongoPort() string     // return the mongo server instance port
}
