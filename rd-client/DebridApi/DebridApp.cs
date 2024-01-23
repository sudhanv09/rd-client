using Aria2NET;
using rd_client.Models;
using Spectre.Console;

namespace rd_client.DebridApi;

public class DebridApp
{
    private readonly RdClient _rdClient;
    private readonly Aria2NetClient _aria;
    
    public DebridApp()
    {
        _rdClient = new RdClient();
        _aria = new Aria2NetClient("http://localhost:6800/jsonrpc", "");
    }
    public async Task DownloadTorrent(string link, bool watch)
    {
        var ddl = await _rdClient.RdUnrestrictLink(link);
    }

    public async Task AddMagnet(string magnet)
    {
        var addMagnet = await _rdClient.RdAddMagnet(magnet);
        await _rdClient.RdFileSelect(addMagnet.Id);

        AnsiConsole.Markup("Generating Download Link\n");
        
        var torrId = await _rdClient.RdTorrentbyId(addMagnet.Id);
        if (torrId is null) throw new Exception("Something went wrong");
        if (torrId.Status != "downloaded") throw new Exception("Torrent not ready yet");
        
        AnsiConsole.Markup($"Found {torrId.Files.Count} files. Adding to aria2\n");
        foreach (var download in torrId.Links )
        {
            var unrestrict = await _rdClient.RdUnrestrictLink(download);
            await _aria.AddUriAsync([unrestrict.Download], new Dictionary<string, object>
            {
                { "dir", "/hdd/media/aria" }
            });
        }
    }
}