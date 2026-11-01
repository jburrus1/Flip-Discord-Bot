using Backend.Data;
using Backend.Models;
using Microsoft.EntityFrameworkCore;

namespace Backend.Services;

public class ChannelsService
{
    private readonly DiscordDbContext _dbContext;

    public ChannelsService(DiscordDbContext dbContext)
    {
        _dbContext = dbContext;
    }
    public async Task<DiscordChannel> CreateAsync(DiscordChannel channel)
    {
        var existing = await _dbContext.Channels.FirstOrDefaultAsync(c => c.Id == channel.Id);
        if (existing != null)
            return existing;

        await _dbContext.Channels.AddAsync(channel);
        await _dbContext.SaveChangesAsync();

        return channel;
    }

    public async Task<DiscordChannel?> GetByIdAsync(string id)
    {
        return await _dbContext.Channels.FindAsync(id);
    }
    
    public async Task<List<DiscordChannel>> GetByServerIdAsync(string id)
    {
        return await _dbContext.Channels.Where(c => c.ServerId == id).ToListAsync();
    }

    public async Task<bool> DeleteAsync(string id)
    {
        var channel = await _dbContext.Channels.FindAsync(id);
        if (channel != null)
        {
            _dbContext.Channels.Remove(channel);
            await _dbContext.SaveChangesAsync();

            return true;
        }

        return false;
    }
}