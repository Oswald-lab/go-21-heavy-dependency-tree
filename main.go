package main

import (
	"fmt"

	_ "github.com/Masterminds/semver/v3"
	_ "github.com/bits-and-blooms/bitset"
	_ "github.com/cespare/xxhash/v2"
	_ "github.com/davecgh/go-spew/spew"
	_ "github.com/dustin/go-humanize"
	_ "github.com/fatih/structs"
	_ "github.com/go-chi/chi/v5"
	_ "github.com/go-chi/cors"
	_ "github.com/gofrs/uuid"
	_ "github.com/google/go-cmp/cmp"
	_ "github.com/google/uuid"
	_ "github.com/gorilla/handlers"
	_ "github.com/gorilla/mux"
	_ "github.com/gorilla/schema"
	_ "github.com/gorilla/sessions"
	_ "github.com/gorilla/websocket"
	_ "github.com/hashicorp/go-multierror"
	_ "github.com/hashicorp/go-version"
	_ "github.com/joho/godotenv"
	_ "github.com/julienschmidt/httprouter"
	_ "github.com/justinas/alice"
	_ "github.com/mitchellh/mapstructure"
	_ "github.com/oklog/ulid/v2"
	_ "github.com/olekukonko/tablewriter"
	_ "github.com/patrickmn/go-cache"
	_ "github.com/pkg/errors"
	_ "github.com/robfig/cron/v3"
	_ "github.com/rs/cors"
	_ "github.com/segmentio/ksuid"
	_ "github.com/shopspring/decimal"
	_ "github.com/spf13/pflag"
	_ "github.com/tidwall/gjson"
)

func main() {
	fmt.Println("go-heavy-deps: 32 direct dependencies loaded")
}
