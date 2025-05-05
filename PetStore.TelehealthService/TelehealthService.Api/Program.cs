using Microsoft.OpenApi.Models;
using Telehealth.AI;
using TelehealthService.Core.Abstractions;
using TelehealthService.JitsiIntegration;
using TelehealthService.JitsiIntegration.Services;

var builder = WebApplication.CreateBuilder(args);

builder.Services.AddControllers();

builder.Services.AddEndpointsApiExplorer();
builder.Services.AddSwaggerGen(c =>
{
    c.SwaggerDoc("v1", new OpenApiInfo { Title = "Telehealth Service", Version = "v1" });
});

builder.Services.AddCors(options =>
{
    var allowedOrigins = builder.Configuration.GetSection("CorsSettings:AllowedOrigins").Get<string[]>()!;
    options.AddPolicy("AllowSpecificOrigin", policy =>
    {
        policy.WithOrigins(allowedOrigins)
            .AllowAnyHeader()
            .AllowAnyMethod()
            .AllowCredentials();
    });
});

builder.Services.Configure<JitsiOptions>(builder.Configuration.GetSection("Jitsi"));
builder.Services.AddScoped<IJitsiService, JitsiService>();

builder.Services.Configure<AiServiceSettings>(builder.Configuration.GetSection("GigaChat"));
builder.Services.AddHttpClient("GigaChatClient", client => 
{
    client.BaseAddress = new Uri("https://ngw.devices.sberbank.ru:9443");
}).ConfigurePrimaryHttpMessageHandler(() => new HttpClientHandler
{
    ServerCertificateCustomValidationCallback = (m, c, ch, e) => true
});
builder.Services.AddScoped<IAiAssistantService, AiAssistantService>();

builder.Services.AddSignalR(options => {
    options.EnableDetailedErrors = true;
    options.ClientTimeoutInterval = TimeSpan.FromSeconds(30);
    options.KeepAliveInterval = TimeSpan.FromSeconds(15);
});

var app = builder.Build();

app.UseStaticFiles();
app.MapControllers();
app.UseHttpsRedirection();

app.UseCors("AllowSpecificOrigin");

app.UseRouting();
app.UseAuthorization();

if (app.Environment.IsDevelopment())
{
    app.UseSwagger();
    app.UseSwaggerUI(c => c.SwaggerEndpoint("/swagger/v1/swagger.json", 
        "Telehealth Service v1"));
}

app.UseEndpoints(endpoints =>
{
    endpoints.MapHub<ChatHub>("/chatHub");
    endpoints.MapControllers();
});

app.Run();