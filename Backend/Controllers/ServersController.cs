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
}