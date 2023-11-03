using System.Diagnostics;

namespace rd_client.Lib;

public class ClipboardWatcher
{
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

    public async IAsyncEnumerable<string> PollClipboard()
    {
        string lastClipText = "";
        while (true)
        {
            string currentText = await GetClipboard();
            if (currentText != lastClipText)
            {
                yield return currentText;
                lastClipText = currentText;
            }
            await Task.Delay(1000);
        }
    }
}