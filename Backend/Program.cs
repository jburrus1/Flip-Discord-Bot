using Backend.Data;
using Backend.Services;
using Microsoft.EntityFrameworkCore;

var builder = WebApplication.CreateBuilder(args);
var connectionString = builder.Configuration.GetConnectionString("DiscordDb");
// Add services to the container.

builder.Services.AddEndpointsApiExplorer();
builder.Services.AddSwaggerGen();
builder.Services.AddControllers();
builder.Services.AddDbContext<DiscordDbContext>(options =>
    options.UseNpgsql(connectionString));
// Learn more about configuring OpenAPI at https://aka.ms/aspnet/openapi
builder.Services.AddOpenApi();
builder.Services.AddHealthChecks();

builder.Services.AddScoped<MessagesService>();
builder.Services.AddScoped<UsersService>();
builder.Services.AddScoped<ChannelsService>();
builder.Services.AddScoped<ServersService>();
builder.Services.AddCors(options =>
{
    options.AddPolicy("BlazorFrontend", policy =>
    {
        policy.WithOrigins("http://localhost:5052")
            .AllowAnyHeader()
            .AllowAnyMethod();
    });
});
var app = builder.Build();

// Configure the HTTP request pipeline.
if (app.Environment.IsDevelopment())
{
    app.MapOpenApi();
    app.UseSwagger();
    app.UseSwaggerUI();
}

using var scope = app.Services.CreateScope();
var dbContext = scope.ServiceProvider.GetRequiredService<DiscordDbContext>();
await dbContext.Database.MigrateAsync();

// app.UseHttpsRedirection();
app.UseCors("BlazorFrontend");
app.UseAuthorization();

app.MapControllers();
app.MapHealthChecks("/health");

app.Run();