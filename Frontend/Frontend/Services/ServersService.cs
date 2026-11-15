using System.Net.Http.Json;
using Frontend.Models;

namespace Frontend.Services;

public class ServersService
{
    private readonly HttpClient _http;

    public ServersService(HttpClient http)
    {
        _http = http;
    }
    
    public async Task<List<DiscordServer>> GetServersAsync()
    {
        return await _http.GetFromJsonAsync<List<DiscordServer>>("api/Servers")
               ?? new List<DiscordServer>();
    }
}