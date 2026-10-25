using Backend.Data;
using Backend.Models;
using Microsoft.EntityFrameworkCore;

namespace Backend.Services;

public class MessagesService
{
    private readonly DiscordDbContext _dbContext;

    public MessagesService(DiscordDbContext dbContext)
    {
        _dbContext = dbContext;
    }
    
    public async Task AddAsync(DiscordMessage message)
    {
        await _dbContext.Messages.AddAsync(message);
        await _dbContext.SaveChangesAsync();
    }
    
    public async Task<bool> DeleteAsync(string id)
    {
        var message = await _dbContext.Messages.FindAsync(id);
        if (message != null)
        {
            _dbContext.Messages.Remove(message);
            await _dbContext.SaveChangesAsync();
            return true;
        }
        return false;
    }
    
    public async Task<bool> DeleteAllAsync()
    {
        _dbContext.Messages.RemoveRange(_dbContext.Messages);
        await _dbContext.SaveChangesAsync();
        return true;
    }

    // Probably not needed. Maybe if a message is edited?
    public async Task<bool> UpdateAsync(string id, DiscordMessage message)
    {
        var existingMessage = await _dbContext.Messages.FindAsync(id);
        if (existingMessage != null)
        {
            existingMessage.Content = message.Content;
            existingMessage.AuthorId = message.AuthorId;
            existingMessage.ChannelId = message.ChannelId;
            await _dbContext.SaveChangesAsync();
            return true;
        }
        return false;
    }

    public async Task<DiscordMessage?> GetByIdAsync(string id)
    {
        return await _dbContext.Messages.FindAsync(id);
    }
    
    public async Task<List<DiscordMessage>> GetByUserIdAsync(string userId)
    {
        return await _dbContext.Messages.Where(m => m.AuthorId == userId).ToListAsync();
    }
    
    public async Task<List<DiscordMessage>> GetAllAsync()
    {
        return await _dbContext.Messages.ToListAsync();
    }
}