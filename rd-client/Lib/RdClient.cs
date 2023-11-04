using System.Collections.Specialized;
using System.Net.Http.Headers;
using System.Text;
using System.Text.Json;
using rd_client.Models;
using System.Web;

namespace rd_client.Lib;

public class RdClient : IRdClient
{
    private readonly HttpClient _client;
    private string _rdApi = "https://api.real-debrid.com/rest/1.0/";
    private NameValueCollection _query = HttpUtility.ParseQueryString(String.Empty);
    public RdClient(HttpClient client, string apiKey)
    {
        _client = client;
        _client.DefaultRequestHeaders.Authorization = new AuthenticationHeaderValue("Bearer", apiKey);
    }
    public async Task<RdUser> RdGetUser()
    {
        string url = _rdApi + "user";
        var req = await _client.GetAsync(url);
        Stream request = await req.Content.ReadAsStreamAsync();

        var result = await JsonSerializer.DeserializeAsync<RdUser>(request);
        return result;
    }

    public async Task<RdTorrents> RdGetUserTorrents()
    {
        string url = _rdApi + "torrents";
        var req = await _client.GetAsync(url);
        Stream request = await req.Content.ReadAsStreamAsync();

        var result = await JsonSerializer.DeserializeAsync<RdTorrents>(request);
        return result;
    }

    public async Task<RdMagnetResult> RdAddMagnet(string magnet)
    {
        _query.Add("magnet", magnet);
        var payload = _query.ToString();
        
        string url = _rdApi + "torrents/addMagnet/" + payload;
        
        var req = await _client.GetAsync(url);
        Stream request = await req.Content.ReadAsStreamAsync();

        var result = await JsonSerializer.DeserializeAsync<RdMagnetResult>(request);
        return result;
    }

    public async Task<RdTorrentId> RdTorrentbyId(string id)
    {
        string url = _rdApi + $"torrent/{id}";
        var req = await _client.GetAsync(url);
        Stream request = await req.Content.ReadAsStreamAsync();

        var result = await JsonSerializer.DeserializeAsync<RdTorrentId>(request);
        return result;
    }

    public async Task RdFileSelect(string id)
    {
        var torrFiles = await RdTorrentbyId(id);
        var files = new Helper().RdFileSelector(torrFiles);

        var payload = new StringContent(files, Encoding.UTF8, "application/json");
        var url = _rdApi + $"torrents/selectFiles/{id}";

        await _client.PostAsync(url, payload);
    }

    public async Task<RdUnrestrict> RdUnrestrictLink(string link)
    {
        _query.Add("link", link);
        var payload = _query.ToString();
        
        string url = _rdApi + "unrestrict/link/" + payload;
        var req = await _client.GetAsync(url);
        Stream request = await req.Content.ReadAsStreamAsync();

        var result = await JsonSerializer.DeserializeAsync<RdUnrestrict>(request);
        return result;
    }
}