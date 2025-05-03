FROM mcr.microsoft.com/dotnet/aspnet:9.0 AS base
WORKDIR /app
EXPOSE 80

FROM mcr.microsoft.com/dotnet/sdk:9.0 AS build
WORKDIR /src

COPY ["TelehealthService.Api/TelehealthService.Api.csproj", "TelehealthService.Api/"]
COPY ["TelehealthService.AI/TelehealthService.AI.csproj", "TelehealthService.AI/"]
COPY ["TelehealthService.Chat/TelehealthService.Chat.csproj", "TelehealthService.Chat/"]
COPY ["TelehealthService.Core/TelehealthService.Core.csproj", "TelehealthService.Core/"]
COPY ["TelehealthService.JitsiIntegration/TelehealthService.JitsiIntegration.csproj", "TelehealthService.JitsiIntegration/"]
RUN dotnet restore "TelehealthService.Api/TelehealthService.Api.csproj"
COPY . .
WORKDIR "/src/TelehealthService.Api"
RUN dotnet build "TelehealthService.Api.csproj" -c Release -o /app/build

FROM build AS publish
RUN dotnet publish "TelehealthService.Api.csproj" -c Release -o /app/publish /p:UseAppHost=false

FROM base AS final
WORKDIR /app
COPY --from=publish /app/publish .
ENTRYPOINT ["dotnet", "TelehealthService.Api.dll"]