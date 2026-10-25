using Microsoft.EntityFrameworkCore;

namespace Backend.Data;

public class DiscordDbContext(DbContextOptions<DiscordDbContext> options) : DbContext(options)
{
}