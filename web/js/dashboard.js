const API_BASE = "https://127.0.0.1:4444";

// ✅ Desactivar verificación SSL para desarrollo
fetch = (function(originalFetch) {
    return function(...args) {
        return originalFetch.apply(this, args);
    };
})(fetch);

async function loadDashboard() {
    try {
        // Obtener clientes conectados
        const clientsResp = await fetch(`${API_BASE}/get_connected_clients`, {
            method: 'GET',
            headers: { 'Content-Type': 'application/json' },
        });
        
        if (clientsResp.ok) {
            const clientsData = await clientsResp.json();
            const implantCount = clientsData.connected_clients.length;
            
            document.getElementById('implant-count').textContent = implantCount;
            document.getElementById('status').textContent = implantCount > 0 ? '● Online' : '● Offline';
            document.getElementById('status').className = implantCount > 0 ? 'text-success' : 'text-danger';
            
            // Cargar tabla de implantes
            loadImplantsList(clientsData.connected_clients);
        }
    } catch (error) {
        console.error('Error cargando dashboard:', error);
        document.getElementById('status').textContent = '● Error';
        document.getElementById('status').className = 'text-warning';
    }
}

async function loadImplantsList(clients) {
    const tbody = document.getElementById('implants-list');
    tbody.innerHTML = '';
    
    for (const clientId of clients) {
        const row = document.createElement('tr');
        row.innerHTML = `
            <td><strong>${clientId}</strong></td>
            <td><span class="badge bg-secondary">-</span></td>
            <td><span class="badge bg-secondary">-</span></td>
            <td><span class="badge bg-secondary">-</span></td>
            <td><span class="badge bg-secondary">-</span></td>
            <td><span class="badge bg-success">Active</span></td>
            <td><small class="text-muted">Now</small></td>
            <td>
                <button class="btn btn-sm btn-info" onclick="openTerminal('${clientId}')">Terminal</button>
            </td>
        `;
        tbody.appendChild(row);
    }
    
    if (clients.length === 0) {
        tbody.innerHTML = '<tr><td colspan="8" class="text-center text-muted">No implants connected</td></tr>';
    }
}

function openTerminal(clientId) {
    alert(`Terminal para ${clientId}\n\nFeature coming soon...`);
}

function logout() {
    if (confirm('¿Estás seguro de que quieres cerrar sesión?')) {
        window.location.href = '/login.html';
    }
}

// Actualizar dashboard cada 5 segundos
document.addEventListener('DOMContentLoaded', function() {
    loadDashboard();
    setInterval(loadDashboard, 5000);
});
