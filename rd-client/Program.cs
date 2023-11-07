
using System.Diagnostics;
using Aria2NET;
using rd_client.Lib;
using Spectre.Console;

var client = new HttpClient();
var rdClient = new RdClient(client, "");
var aria = new Aria2NetClient("http://localhost:6800/jsonrpc", "");

async Task<List<string>> DebridInit(string clipboard)
{
    // Add magnet and choose files
    var magnetId = await rdClient.RdAddMagnet(clipboard);
    await rdClient.RdFileSelect(magnetId.Id);
    
    // Get Links from torrent, check if RD has completed download and unrestrict them.
    var fileInfo = await rdClient.RdTorrentbyId(magnetId.Id);
    var unrestrictedLinks = new List<string>();
    
    if (fileInfo.Status == "downloaded")
    {
        foreach (var link in fileInfo.Links)
        {
            var links = await rdClient.RdUnrestrictLink(link);
            unrestrictedLinks.Add(links.Download);
        }
    }
    return unrestrictedLinks;
}

async Task App()
{
    var downloadCount = 0;
    AnsiConsole.Markup("[green] App running. Watching for magnet links...[/]");
    await foreach (var magnet in new ClipboardWatcher().PollClipboard())
    {
        if (new Helper().ValidMagnet(magnet))
        {
            var downloadLink = await DebridInit(magnet);
            var addUri = await aria.AddUriAsync(downloadLink, new Dictionary<string, object>()
            {
                { "dir", "/hdd/media/aria" },
                {"max-concurrent-downloads", "3"},
                {"split", "16"},
                // {"max-connection-per-server", "16"},
                
            }, downloadCount);
            downloadCount += 1;
            
            var active = await aria.TellActiveAsync();
            foreach (var downloads in active)
            {
                var progress = (double)(downloads.CompletedLength / downloads.TotalLength) * 100;
                AnsiConsole.Markup($"[blue] {downloads.Files[0].Path} [yellow]{progress} [green]{downloads.DownloadSpeed} [/]");
            }
        }
    }
}

await App();