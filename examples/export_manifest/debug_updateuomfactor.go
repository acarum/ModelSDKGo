package main
import (
"database/sql"
"fmt"
"os"
"path/filepath"
"strings"
_ "github.com/mattn/go-sqlite3"
"go.mongodb.org/mongo-driver/bson"
"go.mongodb.org/mongo-driver/bson/primitive"
)
func main() {
mprPath := "C:\\Workspaces\\Mendix\\MDUI\\System_Mendix_CLI\\OC EX System.mpr"
db, _ := sql.Open("sqlite3", mprPath)
defer db.Close()
contentsDir := filepath.Join(filepath.Dir(mprPath), "mprcontents")
rows, _ := db.Query("SELECT UnitID FROM Unit")
defer rows.Close()
for rows.Next() {
var unitIDBytes []byte
rows.Scan(&unitIDBytes)
unitID := guidToString(unitIDBytes)
content, err := loadUnitContents(contentsDir, unitID)
if err != nil { continue }
name, ok := content["Name"].(string)
if !ok || name != "UpdateUoMFactor" { continue }
fmt.Printf("\n✅ Found %s\n\n", name)
findExternalActions(content, 0, 10)
break
}
}
func findExternalActions(obj interface{}, depth, maxDepth int) {
if depth > maxDepth { return }
indent := strings.Repeat("  ", depth)
switch v := obj.(type) {
case map[string]interface{}:
if typeVal, ok := v["$Type"].(string); ok && typeVal == "Microflows$ExternalAction" {
fmt.Printf("%s🎯 ExternalAction found:\n", indent)
if app, ok := v["AppName"].(string); ok {
fmt.Printf("%s  AppName: %s\n", indent, app)
}
if cmd, ok := v["CommandName"].(string); ok {
fmt.Printf("%s  CommandName: %s\n", indent, cmd)
}
}
for _, val := range v {
findExternalActions(val, depth+1, maxDepth)
}
case primitive.A:
for _, item := range v {
findExternalActions(item, depth+1, maxDepth)
}
case []interface{}:
for _, item := range v {
findExternalActions(item, depth+1, maxDepth)
}
}
}
func guidToString(guidBytes []byte) string {
if len(guidBytes) != 16 { return "" }
return fmt.Sprintf("%02x%02x%02x%02x-%02x%02x-%02x%02x-%02x%02x-%02x%02x%02x%02x%02x%02x",
guidBytes[3], guidBytes[2], guidBytes[1], guidBytes[0],
guidBytes[5], guidBytes[4], guidBytes[7], guidBytes[6],
guidBytes[8], guidBytes[9], guidBytes[10], guidBytes[11], guidBytes[12], guidBytes[13], guidBytes[14], guidBytes[15])
}
func loadUnitContents(contentsDir, unitID string) (map[string]interface{}, error) {
cleanID := strings.ReplaceAll(unitID, "-", "")
filePath := filepath.Join(contentsDir, cleanID[0:2], cleanID[2:4], unitID+".mxunit")
data, err := os.ReadFile(filePath)
if err != nil { return nil, err }
var content map[string]interface{}
bson.Unmarshal(data, &content)
return content, nil
}
