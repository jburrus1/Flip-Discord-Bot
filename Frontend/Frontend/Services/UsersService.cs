using System.Net.Http.Json;
using Frontend.Models;

public class UsersService
{
    private readonly HttpClient _http;

    public UsersService(HttpClient http)
    {
        _http = http;
    }

    public async Task<List<DiscordUser>> GetUsersAsync()
    {
        return await _http.GetFromJsonAsync<List<DiscordUser>>("api/Users")
               ?? new List<DiscordUser>();
    }
}