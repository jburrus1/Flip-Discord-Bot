using Backend.Models;
using Backend.Services;
using Microsoft.AspNetCore.Mvc;

namespace Backend.Controllers;

[ApiController]
[Route("api/[controller]")]
public class ServersController : ControllerBase
{
    private readonly ServersService _serversService;

    public ServersController(ServersService serversService)
    {
        _serversService = serversService;
    }

    [HttpPost]
    public async Task<ActionResult<DiscordServer>> Create(DiscordServer server)
    {
        await _serversService.CreateAsync(server);
        return CreatedAtAction(nameof(GetById), new { id = server.Id }, server);
    }

    [HttpGet]
    public async Task<ActionResult<List<DiscordServer>>> GetAllServers()
    {
        var servers = await _serversService.GetAllAsync();

        return Ok(servers);
    }

    [HttpGet("{id}")]
    public async Task<ActionResult<DiscordServer>> GetById(string id)
    {
        var server = await _serversService.GetByIdAsync(id);
        if (server == null)
        {
            return NotFound();
        }

        return Ok(server);
    }

    [HttpDelete("{id}")]
    public async Task<ActionResult> Delete(string id)
    {
        await _serversService.DeleteAsync(id);
        return NoContent();
    }
    
    [HttpPost("{serverId}/users")]
    public async Task<ActionResult> AddUser(string serverId, AddServerMemberRequest memberRequest)
    {
        await _serversService.AddUserToServerAsync(serverId, memberRequest.UserId);
        return NoContent();
    }
    [HttpGet("{serverId}/users")]
    public async Task<ActionResult<List<DiscordUser>>> GetUsers(string serverId)
    {
        
        if (!await _serversService.ExistsAsync(serverId))
        {
            return NotFound();
        }
        var users = await _serversService.GetUsersAsync(serverId);
        return Ok(users);
    }
    [HttpGet("{serverId}/messages")]
    public async Task<ActionResult<List<DiscordMessage>>> GetMessages(string serverId)
    {
        
        if (!await _serversService.ExistsAsync(serverId))
        {
            return NotFound();
        }
        var messages = await _serversService.GetMessagesAsync(serverId);
        return Ok(messages);
    }
}