package	small_test

// The log will manage the number of segments

import (
"fmt"
"io"
"io/ioutil"
"os"
"path"
"sort"
"strconv"
"strings"
"sync"
api "github.com/upalchowdhury/dist-service/api/v1"
)



// Log consists of list of segments and pointer to the active segment to append writes to and the directory where we store the segments
 

files, err := ioutil.ReadDir(l.Dir)

var baseoffset []uint64

