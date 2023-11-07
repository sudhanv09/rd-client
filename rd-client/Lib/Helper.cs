using System.Diagnostics;
using System.Text.RegularExpressions;
using Aria2NET;
using rd_client.Models;
using Spectre.Console;

namespace rd_client.Lib;

public class Helper
{
    public bool ValidMagnet(string magnet)
    {
        var regex = @"^magnet:\?xt=urn:btih:[0-9a-fA-F]{40}";
        
        var rg = new Regex(regex);
        var match = rg.Match(magnet);
        if (!match.Success)
        {
            return false;
        }
        return true;
    }

    /// <summary>
    /// Only select video files and subtitles from the torrent files
    /// </summary>
    /// <param name="files"></param>
    /// <returns>String with all the selected file ids</returns>
    public string RdFileSelector(RdTorrentId files)
    {
        var allowedFileTypes = new List<string>() { "mkv", "mp4", "srt" };
        
        // If the video filename contains allowedFileTypes, get the Id
        List<string> idx = files.Files
            .Where(file => allowedFileTypes.Any(t => file.Path.Contains(t)))
            .Select(i => i.Id.ToString())
            .ToList();
        
        return string.Join(",", idx);
    }

    public async Task MonitorDownloads(Aria2NetClient aria)
    {
        var downloadsComplete = false;
        while (!downloadsComplete)
        {
            var pollAria = await aria.TellActiveAsync();
            foreach (var downloads in pollAria)  
            {
                if (downloads.Status == "complete")
                {
                    downloadsComplete = true;
                }
                // TODO Add a progress bar
            }
        }
    }
}