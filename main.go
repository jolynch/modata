package main

import (
    "flag"
    "modata"
    "fmt"
    "time"
)

func main() {
    isBlock := flag.Bool("block", true, "Run the Block server")
    isReplication := flag.Bool("replication", false, "Run the replication server")

    blockName := flag.String("block-server",
                             "localhost:1234",
                             "Name of block server")
    replicationName := flag.String("replication-server",
                                   "localhost:1337",
                                   "Name of replication server")
    bootstrapName := flag.String("bootstrap",
                                 "",
                                 "Bootstrap server and port to contact")

    flag.Parse()

    fmt.Printf("Block: %v at %v\n", *isBlock, *blockName)
    fmt.Printf("Replication: %v at %v\n", *isReplication, *replicationName)

    // Block Server
    var bs *modata.BlockServer
    if (*isBlock) {
        fmt.Println("Starting block server!")
        bs = modata.StartBlockServer(*blockName)
        // Ping the bootstrap
        fmt.Printf("Pinging the bootstrap ... %s\n", "http://" + *bootstrapName + "/ping")
        _, _, contact := modata.JsonGet("http://" + *bootstrapName + "/ping", bs.Contact())
        bs.UpdateContact(contact)
    }

    // Replication Server
    var rs *modata.ReplicationServer
    if (*isReplication) {
        fmt.Println("Starting replication server")
        rs = modata.StartReplicationServer(*replicationName, bs)
    }

    fmt.Println(bs)
    fmt.Println(rs)

    // Spin loop foreverz
    for *isBlock || *isReplication {
        // Yield to other threads, usually the server threads
        time.Sleep(100 * time.Second)
    }
}
