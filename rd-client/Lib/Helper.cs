using System.Diagnostics;
using System.Text.RegularExpressions;
using rd_client.Models;

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

    public string RdFileSelector(RdTorrentId files)
    {
        var allowedFileTypes = new List<string>() { "mkv", "mp4", "srt" };
        var fileIds = new List<string>();

        foreach (var file in files.Files)
        {
            var idx = allowedFileTypes.Any(f => file.Path.Contains(f)).ToString();
            fileIds.Append(idx);
        }
        return string.Join(",", fileIds);
    }

    public async Task<string> GetClipboard()
    {
        ProcessStartInfo startInfo = new()
        {
            FileName = "wl-paste",
            CreateNoWindow = true,
            RedirectStandardError = true,
            RedirectStandardOutput = true
        };

        var proc = Process.Start(startInfo);
        ArgumentNullException.ThrowIfNull(proc);
        
        string clipboard = await proc.StandardOutput.ReadToEndAsync();
        await proc.WaitForExitAsync();

        return clipboard;
    }

    public async Task<string> PollClipboard()
    {
        string lastClipText = "";
        while (true)
        {
            string currentText = await GetClipboard();
            if (currentText != lastClipText)
            {
                lastClipText = currentText;
                return currentText;
            }
            Thread.Sleep(1000);
        }
    }

}