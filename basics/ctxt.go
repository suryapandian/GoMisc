type Job struct {
    ID int
    Context context.Context
    Cancel context.CancelFunc
}
var jobs = make(map[int] Job)

func worker(ctx context.Context, index int) {
   fmt.Printf("starting job with id %d\n", index)
   <-ctx.Done()
}

func main() {
 var err error

 id := 0
 r := gin.Default()

 r.POST("/start", func(c *gin.Context) {
    var job Job
    err := json.NewDecoder(c.Request.Body).Decode(&job)
    if err != nil{ fmt.Println(err)}
    ctx, cancel := context.WithCancel(context.Background())
    job.ID = id
    job.Context = ctx
    job.Cancel = cancel
    jobs[job.ID] = job
    c.JSON(http.StatusOK, gin.H{"message": "job received"})
    go worker(ctx, job.ID)
    id ++
 })

 r.GET("/cancel/:id", func(c *gin.Context) {
    id := c.Param("id")
    idInt, err := strconv.Atoi(id)
    if err != nil {fmt.Println(err)}
    jobs[idInt].Cancel()
 })

  err = endless.ListenAndServe(":8080", r)
  if err != nil{ fmt.Printf("Could not run server : %v", err.Error())}
 }
