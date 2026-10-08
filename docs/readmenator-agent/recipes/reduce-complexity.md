# Recipe: Reduce File Complexity

Target hotspot: `app.py`
(complexity 1.0, centrality 0.7)

1. Read dependents: `grep -n 'app.py' readmenator-agent/ARCHITECTURE*.md`
2. Extract functions/classes into new files in the same subsystem
3. Update imports
4. Regenerate: `readmenator .`
