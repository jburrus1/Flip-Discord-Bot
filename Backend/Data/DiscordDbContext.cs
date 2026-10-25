using Backend.Models;
using Microsoft.EntityFrameworkCore;

namespace Backend.Data;

public class DiscordDbContext : DbContext
{
    public DiscordDbContext(
        DbContextOptions<DiscordDbContext> options)
        : base(options)
    {
    }

    public DbSet<DiscordUser> Users => Set<DiscordUser>();
    public DbSet<DiscordMessage> Messages => Set<DiscordMessage>();
    

    protected override void OnModelCreating(ModelBuilder modelBuilder)
    {
        modelBuilder.Entity<DiscordMessage>()
            .HasOne(m => m.Author)
            .WithMany()
            .HasForeignKey(m => m.AuthorId)
            .OnDelete(DeleteBehavior.Restrict);
    }
}