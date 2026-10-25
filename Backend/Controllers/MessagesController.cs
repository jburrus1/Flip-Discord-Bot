using Backend.Models;
using Backend.Services;
using Microsoft.AspNetCore.Mvc;

namespace Backend.Controllers;

[ApiController]
[Route("api/[controller]")]
public class MessagesController : ControllerBase
{
    private readonly MessagesService _messagesService;

    public MessagesController(MessagesService messagesService)
    {
        _messagesService = messagesService;
    }

    [HttpPost]
    public async Task<ActionResult<DiscordMessage>> Create(DiscordMessage message)
    {
        await _messagesService.AddAsync(message);
        return CreatedAtAction(nameof(GetById), new { id = message.Id }, message);
    }

    [HttpGet("user/{userId}")]
    public async Task<ActionResult<List<DiscordMessage>>> GetByUserId(string userId)
    {
        var messages = await _messagesService.GetByUserIdAsync(userId);
        return Ok(messages);
    }

    [HttpGet("{id}")]
    public async Task<ActionResult<DiscordMessage>> GetById(string id)
    {
        var message = await _messagesService.GetByIdAsync(id);
        if (message == null)
        {
            return NotFound();
        }

        return Ok(message);
    }

    [HttpDelete("{id}")]
    public async Task<ActionResult> Delete(string id)
    {
        await _messagesService.DeleteAsync(id);
        return NoContent();
    }

    [HttpDelete]
    public async Task<ActionResult> DeleteAll()
    {
        await _messagesService.DeleteAllAsync();
        return NoContent();
    }
}