using System.Collections.Specialized;
using System.Net.Http.Headers;
using System.Text;
using System.Text.Json;
using System.Web;
using rd_client.Lib;
using rd_client.Models;

namespace rd_client.DebridApi;

public class RdClient : IRdClient
{
    private readonly HttpClient _client;
    private string _rdApi = "https://api.real-debrid.com/rest/1.0/";
    private NameValueCollection _query = HttpUtility.ParseQueryString(string.Empty);
    public RdClient()
    {
        _client = new HttpClient();
        var apiKey = Environment.GetEnvironmentVariable("rd_api");
        _client.DefaultRequestHeaders.Authorization = new AuthenticationHeaderValue("Bearer", apiKey);
    }
    public async Task<RdUser> RdGetUser()
    {
        var url = _rdApi + "user";
        var req = await _client.GetAsync(url);
        var request = await req.Content.ReadAsStreamAsync();

        var result = await JsonSerializer.DeserializeAsync<RdUser>(request);
        return result;
    }

    public async Task<List<RdTorrents>> RdGetUserTorrents()
    {
        var url = _rdApi + "torrents";
        var req = await _client.GetAsync(url);
        var request = await req.Content.ReadAsStreamAsync();

        var result = await JsonSerializer.DeserializeAsync<List<RdTorrents>>(request);
        return result;
    }

    public async Task<RdMagnetResult> RdAddMagnet(string magnet)
    {
        _query.Add("magnet", magnet);
        var payload = new StringContent(_query.ToString(), Encoding.UTF8, "application/json");
        var url = _rdApi + "torrents/addMagnet/";
        
        var req = await _client.PostAsync(url, payload);
        var request = await req.Content.ReadAsStreamAsync();

        var result = await JsonSerializer.DeserializeAsync<RdMagnetResult>(request);
        return result;
    }

    public async Task<RdTorrentId> RdTorrentbyId(string id)
    {
        var url = _rdApi + $"torrents/info/{id}";
        var req = await _client.GetAsync(url);
        var request = await req.Content.ReadAsStreamAsync();

        var result = await JsonSerializer.DeserializeAsync<RdTorrentId>(request);
        return result;
    }

    public async Task RdFileSelect(string id)
    {
        var torrFiles = await RdTorrentbyId(id);
        var files = new Helper().RdFileSelector(torrFiles);
        _query.Add("files", files);
        
        var payload = new StringContent(_query.ToString(), Encoding.UTF8, "application/json");
        var url = _rdApi + $"torrents/selectFiles/{id}";

        await _client.PostAsync(url, payload);
    }

    public async Task<RdUnrestrict> RdUnrestrictLink(string link)
    {
        _query.Add("link", link);
        var payload = new StringContent(_query.ToString(), Encoding.UTF8, "application/json");
        
        var url = _rdApi + "unrestrict/link/";
        var req = await _client.PostAsync(url, payload);
        var request = await req.Content.ReadAsStreamAsync();

        var result = await JsonSerializer.DeserializeAsync<RdUnrestrict>(request);
        return result;
    }
}