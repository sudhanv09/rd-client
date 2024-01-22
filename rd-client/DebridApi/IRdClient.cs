using rd_client.Models;

namespace rd_client.DebridApi;

public interface IRdClient
{
    Task<RdUser> RdGetUser();
    Task<List<RdTorrents>> RdGetUserTorrents();
    Task<RdMagnetResult> RdAddMagnet(string magnet);
    Task<RdTorrentId> RdTorrentbyId(string id);
    Task RdFileSelect(string id);
    Task<RdUnrestrict> RdUnrestrictLink(string link);
}