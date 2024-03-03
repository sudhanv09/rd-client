using rd_client.Cli;
using Spectre.Console.Cli;

DotNetEnv.Env.TraversePath().Load();

var app = new CommandApp();
app.Configure(conf =>
{
    conf.AddCommand<AddMagnetCmd>("magnet")
        .WithDescription("Add magnet links");
    
    conf.AddCommand<AddDebridCmd>("debrid")
        .WithDescription("Add debrid links")
        .WithExample(["debrid", "https://real-debrid.com/d/IIxxx...", "--download"]);
    
    conf.AddCommand<GetTorrentsCmd>("list")
        .WithDescription("Get all torrents from real debrid");
});

app.Run(args);






