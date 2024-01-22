using rd_client.DebridApi;
using rd_client.Models;
using Spectre.Console;
using Spectre.Console.Cli;

namespace rd_client.Lib;

public sealed class AddMagnetSetting : CommandSettings
{
    [CommandArgument(0, "<Magnet Link>")]
    public string Magnet { get; set; }
    [CommandOption("-d|--download")]
    public bool? Download { get; set; }
    [CommandOption("-w|--watch")]
    public bool? Stream { get; set; }
}
public sealed class AddDebridSetting : CommandSettings
{
    [CommandArgument(0, "<Debrid Link>")]
    public string Debrid { get; set; }
    [CommandOption("-d|--download")]
    public bool? Download { get; set; }
    [CommandOption("-w|--watch")]
    public bool? Stream { get; set; }
}

public sealed class GetTorrentsSetting : CommandSettings
{
    public string List { get; set; }
}

public class AddMagnetCmd : Command<AddMagnetSetting>
{
    public override int Execute(CommandContext ctx, AddMagnetSetting cmd)
    {
        var helper = new Helper();
        var magnet = helper.ValidMagnet(cmd.Magnet);
        if (!magnet) throw new Exception("Magnet not valid");
        
        AnsiConsole.Markup("Adding magnet...");
        var task = Task.Run(async () =>
        {
            await new DebridApp().AddMagnet(cmd.Magnet);
            
        });
        task.Wait();
        
        return 0;
    }
}

public class AddDebridCmd : Command<AddDebridSetting>
{
    public override int Execute(CommandContext ctx, AddDebridSetting cmd)
    {
        return 0;
    }
}

public class GetTorrentsCmd : Command<GetTorrentsSetting>
{
    public override int Execute(CommandContext ctx, GetTorrentsSetting cmd)
    {
        return 0;
    }
}